package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"

	c "github.com/axyut/playgo/internal/config"
	"github.com/axyut/playgo/internal/tui"

	"github.com/ebitengine/oto/v3"

	// keypress listener
	"github.com/mattn/go-tty"
)

var Notifications []string

// Prepare an Oto context (creating context for single time)
var op = &oto.NewContextOptions{
	SampleRate:   44100,
	ChannelCount: 2,
	Format:       oto.FormatSignedInt16LE,
}
var otoCtx, readyChan, otoErr = oto.NewContext(op)
var playlist []string
var playedList []string

// var favorites []string
// var timer int
var completedPlaylist int

//type pos [2]int

type Player struct {
	Music       *oto.Player
	UserSetting c.UserSetting
	File        *os.File
	Song        int
}

var UserSetting = c.UserSetting{
	Shuffle:        true,
	RepeatSong:     false, // if Shuffle false it's no use for RepeatSong to be true
	RepeatPlaylist: true,
}

var songs = c.Activelist{
	PrevSong:    -1,
	CurrentSong: 0,
	NextSong:    1,
}

var flags = c.Flag{
	Help: "h",
	Test: "t",
}

func shufflePlaylist(playlist *[]string) {
	list := *playlist
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}

func serializePlaylist(playlist *[]string) {
	// just doing addFolder for now which doesn't cover when individual files opened in command $playgo a.mp3 b.mp3
	addFolder(".", playlist)
}

func addFolder(path string, playlist *[]string) error {
	fileInfos, err := os.ReadDir(path)
	if err != nil {
		log.Println("Couldn't Read from Current Directory!")
		return err
	}
	for _, file := range fileInfos {
		ext := filepath.Ext(file.Name())

		if ext == ".mp3" {
			// path, _ := filepath.Abs(filepath.Dir(file.Name()))
			*playlist = append(*playlist, file.Name())
		}
	}

	if len(*playlist) == 0 {
		fmt.Println(c.Usage)
		os.Exit(0)
	}
	return nil
}

func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func handleInterrupt() {
	tui.HideCursor()

	// handle CTRL C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			tui.DisplayStats(playlist, playedList, completedPlaylist)
		}
	}()
}

func listenForKey() {
	tty, err := tty.Open()
	if err != nil {
		panic(err)
	}
	defer tty.Close()
	// listen for keypress
	for {
		if char, err := tty.ReadRune(); err == nil {
			// str := string(char)
			// fmt.Println(str)
			switch char {
			case 'a': // prev Music
				notify("<< Previous")
			case 'q': // seek left
				notify("<< 10s ")
			case 'e': // seek right
				notify(">> 10s")
			case 'd': // next Music
				notify(">> Next")
			case 'p': // play/pause
				notify("|> Paused")
				notify("|| Playing")
			case 'w': // volume up
				notify("++ VOL")
			case 's': // volume down
				notify("-- VOL")
			case 'y': // shuffle
				toogleSetting('y', &playlist, &UserSetting)
			case 't': // repeat playlist
				toogleSetting('t', &playlist, &UserSetting)
			case 'r': // repeat Song
				toogleSetting('r', &playlist, &UserSetting)
			case 'x':
				tui.DisplayStats(playlist, playedList, completedPlaylist)
			}
		}
	}
}

func toogleSetting(str rune, list *[]string, UserSetting *c.UserSetting) {
	suf, repS, repP := UserSetting.Shuffle, UserSetting.RepeatSong, UserSetting.RepeatPlaylist
	switch str {
	case 't':
		repP = !repP
	case 'r':
		repS = !repS
	case 'y':
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
	*UserSetting = c.UserSetting{
		Shuffle:        suf,
		RepeatSong:     repS,
		RepeatPlaylist: repP,
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

// func inPlaylist(songNum int, playlist []string) bool {
// 	for _, v := range playedList {
// 		if playlist[songNum] == v {
// 			return true
// 		}
// 	}
// 	return false
// }

func notify(str string) {
	Notifications = append([]string{str}, Notifications...)
}

func getSong(i int, playlist *[]string, UserSetting c.UserSetting) *c.Activelist {
	var prevSong, curSong, nextSong int
	prevSong = i - 1
	curSong = i
	if len(*playlist) == i+1 {
		nextSong = i
		if UserSetting.RepeatPlaylist {
			nextSong = 0
		}
	} else {
		nextSong = i + 1
	}
	songs = c.Activelist{
		PrevSong:    prevSong,
		CurrentSong: curSong,
		NextSong:    nextSong,
	}
	return &songs
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
