package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
)

func addFolder() {
	fileInfos, err := os.ReadDir(".")
	if err != nil {
		log.Fatal("Couldn't Read from Current Directory!\n")
	}

	for _, file := range fileInfos {
		ext := filepath.Ext(file.Name())
		if ext == ".mp3" {
			playlist = append(playlist, file.Name())
		}
	}
}
func Remove(slice []int, s int) []int {
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
func UniqSong() (songNum int) {

	for {
		songNum = rand.Intn(len(playlist))
		if !UserSetting.RepeatSong {
			if inPlaylist(songNum) {
				continue
			}
		}
		break
	}

	return songNum
}

func inPlaylist(songNum int) bool {
	for _, v := range playedList {
		if playlist[songNum] == v {
			return true
		}
	}
	return false
}

func notify(str string) {
	notifications = append([]string{str}, notifications...)
}
func stripString(str string) string {
	maxX, _ := termSize()
	strip := maxX / 3
	len := len(str)
	if len > strip {
		return str[:strip] + "..."
	}
	return str
}
func getSong(i int) (songNum int) {
	if UserSetting.Shuffle {
		songNum = UniqSong()
	} else {
		songNum = i
	}
	return songNum
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
