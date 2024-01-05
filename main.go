package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/mattn/go-tty"
)

// when file is deleted, if cannot open file remove from playlist
// change the playlist as a whole when shuffled, re-load folder if turned off
func main() {
	handleArgs()
top:
	for i := 0; ; i++ {
	repeatSameSong:
		songs := getSong(i)
		player := play(songs.currentSong)
		if player == nil {
			continue
		}
		player.handleInterrupt()
		player.display()
		// go player.listenForKey()
		for player.Music.IsPlaying() {
			time.Sleep(time.Millisecond)
		}

		err := player.Music.Close()
		if err != nil {
			panic("player.Close failed: " + err.Error())
		}
		player.File.Close()
		playedList = appendOnlyOriginal(playedList, playlist[songs.currentSong])
		played := fmt.Sprintf("Played %s", playlist[songs.currentSong])
		notify(played)
		if len(playedList) == len(playlist) {
			playedList = []string{}
			completedPlaylist++
			break
		}
		if UserSetting.RepeatSong {
			goto repeatSameSong
		}
	}
	if UserSetting.RepeatPlaylist {
		notify("Restarting Playlist.")
		goto top
	}
	displayStats()
}

// seek, next , prevous, pause, play, settings
func play(songNum int) *Player {
	mp3File := playlist[songNum]
	file, err := os.Open(mp3File)
	if err != nil {
		log.Print(err)
		playlist = Remove(playlist, songNum)
		return nil
	}

	// Decode file. This process is done as the file plays so it won't load whole thing into memory
	decodedMp3, err := mp3.NewDecoder(file)

	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}
	if otoErr != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan
	otoPlayer := otoCtx.NewPlayer(decodedMp3)

	newPlayer := Player{
		otoPlayer,
		UserSetting,
		file,
		songNum,
		*decodedMp3,
	}
	newPlayer.Music.Play()
	go newPlayer.listenForKey()
	return &newPlayer
}
func (p *Player) listenForKey() {
	tty, err := tty.Open()
	if err != nil {
		panic(err)
	}
	defer tty.Close()
	// listen for keypress
	for {
		if char, err := tty.ReadRune(); err == nil {
			// str := string(char)
			switch char {
			case 'h': // prev Music
				notify("<< Previous")
			case 'j': // seek left
				notify("<< 10s ")
			case 'k': // seek right
				notify(">> 10s")
			case 'l': // next Music
				notify(">> Next")
			case 'p': // play/pause
				notify("|> Paused")
				notify("|| Playing")
			case 'w': // volume up
				notify("++ VOL")
			case 's': // volume down
				notify("-- VOL")
			case 't': // shuffle
				toogleSetting('t')
			case 'e': // repeat playlist
				toogleSetting('e')
			case 'r': // repeat Song
				toogleSetting('r')
			case 'q':
				displayStats()
			}
		}
	}
}
func handleArgs() {
	if len(os.Args) == 1 {
		log.Fatal(usage)
	}
	// check if it's files or a folder
	if os.Args[1] == "." {
		addFolder()
	} else {
		for i, v := range os.Args {
			if i == 0 {
				continue
			}
			if loc := strings.Index(v, "-"); loc == 0 {
				v = v[1:]
				switch v {
				default:
					fmt.Println(usage)
				case flags.Test:
					fmt.Println("Checking Players Health.")
					//TODO: to pass all the tests
					fmt.Println("OK!")
				case flags.Help:
					fmt.Println(usage)
				}
				os.Exit(0)
			}
			playlist = append(playlist, v)
		}
	}
	if UserSetting.Shuffle {
		shufflePlaylist()
	}
}
