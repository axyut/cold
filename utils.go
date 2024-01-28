package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
)

var Notifications []string

func shufflePlaylist(playlist []string) {
	rand.Shuffle(len(playlist), func(i, j int) {
		playlist[i], playlist[j] = playlist[j], playlist[i]
	})
}
func serializePlaylist(playlist []string) {
	// just doing addFolder for now which doesn't cover when individual files opened in command $playgo a.mp3 b.mp3
	addFolder(".", playlist)
}
func addFolder(path string, playlist []string) error {
	fileInfos, err := os.ReadDir(path)
	if err != nil {
		log.Println("Couldn't Read from Current Directory!\n")
		return err
	}
	for _, file := range fileInfos {
		ext := filepath.Ext(file.Name())

		if ext == ".mp3" {
			// path, _ := filepath.Abs(filepath.Dir(file.Name()))
			playlist = append(playlist, file.Name())
		}
	}

	if len(playlist) == 0 {
		fmt.Println(usage)
		os.Exit(0)
	}
	return nil
}

func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func (player *Player) handleInterrupt() {
	hideCursor()

	// handle CTRL C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			displayStats()
		}
	}()
}

func toogleSetting(str rune, list []string, UserSetting Setting) {
	suf, repS, repP := UserSetting.Shuffle, UserSetting.RepeatSong, UserSetting.RepeatPlaylist
	switch str {
	case 'e':
		repP = !repP
	case 'r':
		repS = !repS
	case 't':
		// serialize
		suf = !suf
		if suf {
			shufflePlaylist(list)
			notify("Shuffle On.")
		} else {
			serializePlaylist(list)
			notify("Shuffle Off.")
		}
	}
	UserSetting = Setting{
		suf, repS, repP,
	}
}

// func UniqSong() (songNum int) {
//
// 	for {
// 		songNum = rand.Intn(len(playlist))
// 		if !UserSetting.RepeatSong {
// 			if inPlaylist(songNum) {
// 				continue
// 			}
// 		}
// 		break
// 	}
//
// 	return songNum
// }

func inPlaylist(songNum int, playlist []string) bool {
	for _, v := range playedList {
		if playlist[songNum] == v {
			return true
		}
	}
	return false
}

func notify(str string) {
	Notifications = append([]string{str}, Notifications...)
}

func stripString(str string) string {
	maxX, _ := termSize()
	strip := maxX / 3
	len := len(str)
	if isMp3 := strings.ContainsAny(str, "mp3"); isMp3 {
		str = str[:len-3]
		len = len - 3
	}
	if len > strip {
		return str[:strip] + "..."
	}
	return str
}

func getSong(i int, playlist []string, UserSetting Setting) Activelist {
	var prevSong, curSong, nextSong int
	prevSong = i - 1
	curSong = i
	if len(playlist) == i+1 {
		nextSong = i
		if UserSetting.RepeatPlaylist {
			nextSong = 0
		}
	} else {
		nextSong = i + 1
	}
	songs = Activelist{
		prevSong, curSong, nextSong,
	}
	return songs
}

func appendOnlyOriginal(list []string, val string) (originalList []string) {
	for _, v := range list {
		if v == val {
			return list
		}
	}
	originalList = append(list, val)
	return originalList
}
