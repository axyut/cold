package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/axyut/playgo/cmd"
	"github.com/axyut/playgo/internal/utils"
)

var playlist = cmd.Playlist
	"fmt"
var usage = L

func ShufflePlaylist() {
	rand.Shuffle(len(playlist), func(i, j int) {
		playlist[i], playlist[j] = playlist[j], playlist[i]
	})
}
func SerializePlaylist() {
	// just doing addFolder for now which doesn't cover when individual files opened in command $playgo a.mp3 b.mp3
	AddFolder()
}
func AddFolder() {
	playlist = []string{}
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

	if len(playlist) == 0 {
		fmt.Println(usage)
		os.Exit(0)
	}
}
func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
func (player *Player) HandleInterrupt() {
	HideCursor()

	// handle CTRL C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			DisplayStats()
		}
	}()
}
func ToogleSetting(str rune) {
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
			ShufflePlaylist()
			Notify("Shuffle On.")
		} else {
			SerializePlaylist()
			Notify("Shuffle Off.")
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

func InPlaylist(songNum int) bool {
	for _, v := range playedList {
		if playlist[songNum] == v {
			return true
		}
	}
	return false
}

func Notify(str string) {
	notifications = append([]string{str}, notifications...)
}
func StripString(str string) string {
	maxX, _ := TermSize()
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

func GetSong(i int) Activelist {
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

func AppendOnlyOriginal(list []string, val string) (originalList []string) {
	for _, v := range list {
		if v == val {
			return list
		}
	}
	originalList = append(list, val)
	return originalList
}
