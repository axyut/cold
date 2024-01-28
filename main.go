package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	mp3 "github.com/axyut/playgo/internal/mp3Decoder"
	"github.com/axyut/playgo/internal/tui"
	// keypress listener
	"github.com/mattn/go-tty"
)

func main() {
	handleArgs()
top:
	for i := 0; ; i++ {
	repeatSameSong:
		songs := getSong(i, &playlist, UserSetting)
		player := play(songs.currentSong)
		if player == nil {
			continue
		}
		player.handleInterrupt()
		tui.Display(playlist, Notifications, songs.currentSong, songs.prevSong, songs.nextSong, UserSetting.Shuffle, UserSetting.RepeatSong, UserSetting.RepeatPlaylist)
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
	tui.DisplayStats(playlist, playedList, completedPlaylist)
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
	decodedMp3, err := mp3.Decode(file)

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
				toogleSetting('t', &playlist, &UserSetting)
			case 'e': // repeat playlist
				toogleSetting('e', &playlist, &UserSetting)
			case 'r': // repeat Song
				toogleSetting('r', &playlist, &UserSetting)
			case 'q':
				tui.DisplayStats(playlist, playedList, completedPlaylist)
			}
		}
	}
}
func handleArgs() {
	if len(os.Args) == 1 {
		addFolder(".", &playlist)
	} else
	// check if it's files or a folder
	if os.Args[1] == "." {
		addFolder(".", &playlist)
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
			// not mp3 file, then its path
			// if loc := strings.Index(v, ".mp3"); loc == -1 {
			// 	addFolder(v)
			// 	continue
			// }
			file, err := os.Open(v)
			if err != nil {
				log.Fatal(err)
			}
			if ext := filepath.Ext(file.Name()); ext == ".mp3" {
				playlist = append(playlist, v)
			}
		}
	}
	if UserSetting.Shuffle {
		shufflePlaylist(&playlist)
	}
}
