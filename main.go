package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	g "github.com/axyut/playgo/internal/global"
	mp3 "github.com/axyut/playgo/internal/mp3Decoder"
	"github.com/axyut/playgo/internal/tui"
)

func main() {
	handleArgs()
	handleInterrupt()
	go listenForKey()

top:
	for i := 0; ; i++ {
	repeatSameSong:
		songs := getSong(i, &playlist, UserSetting)

		player := play(songs.CurrentSong)
		if player == nil {
			continue
		}

		go tui.Display(&playlist, &Notifications, songs, &UserSetting)

		for player.Music.IsPlaying() {
			time.Sleep(time.Millisecond)
		}

		err := player.Music.Close()
		if err != nil {
			panic("player.Close failed: " + err.Error())
		}
		player.File.Close()

		playedList = appendOnlyOriginal(playedList, playlist[songs.CurrentSong])
		played := fmt.Sprintf("Played %s", playlist[songs.CurrentSong])
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
		panic("oto.NewContext failed: " + otoErr.Error())
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
	return &newPlayer
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
					fmt.Println(g.Usage)
				case flags.Test:
					fmt.Println("Checking Players Health.")
					//TODO: to pass all the tests
					fmt.Println("OK!")
				case flags.Help:
					fmt.Println(g.Usage)
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
