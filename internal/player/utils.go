package player

// import (
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"os/signal"
// 	"path/filepath"

// 	"github.com/axyut/playgo/internal/types"
// 	"github.com/axyut/playgo/pkg/raw"

// 	"github.com/ebitengine/oto/v3"
// 	// keypress listener
// )

// var Notifications []string

// // var playlist []string
// var playedList []string
// var playlist = types.Playlist{}

// // var favorites []string
// // var timer int
// var completedPlaylist int

// //type pos [2]int

// var songs = types.Activelist{
// 	PrevSong:    -1,
// 	CurrentSong: 0,
// 	NextSong:    1,
// }

// // var flags = c.Flag{
// // 	Help: "h",
// // 	Test: "t",
// // }

// func shufflePlaylist(playlist *types.Playlist) {
// 	list := playlist.List
// 	rand.Shuffle(len(list), func(i, j int) {
// 		list[i], list[j] = list[j], list[i]
// 	})
// }

// func serializePlaylist(path string) {
// 	// just doing addFolder for now which doesn't cover when individual files opened in command $playgo a.mp3 b.mp3
// 	if err, _ := addFolder(path); err != nil {
// 		log.Default().Println(err)
// 	}
// }

// func addFolder(path string) (*types.Playlist, error) {
// 	// get full path
// 	if path == "." {
// 		path, _ = filepath.Abs(path)
// 	}
// 	fmt.Println("Adding Folder", path)
// 	fileInfos, err := os.ReadDir(path)
// 	if err != nil {
// 		log.Println("Couldn't Read from ", path, " Directory!")
// 		return &playlist, err
// 	}
// 	for _, file := range fileInfos {
// 		ext := filepath.Ext(file.Name())
// 		if ext == ".mp3" {
// 			// path, _ := filepath.Abs(filepath.Dir(file.Name()))
// 			song := types.Song{
// 				Name: file.Name(),
// 				Path: path,
// 			}

// 			playlist.List = append(playlist.List, song)
// 		}
// 	}
// 	home, _ := os.UserHomeDir()
// 	musicPath := filepath.Join(home + "/Music/")
// 	// fmt.Println("Music Path: ", musicPath)
// 	if len(playlist.List) == 0 && path == musicPath {
// 		fmt.Println(types.Usage)
// 		os.Exit(0)
// 	} else if len(playlist.List) == 0 {
// 		fmt.Println("No Music Files Found in the given path:", path, "Trying ~/Music/")
// 		addFolder(musicPath)
// 	}
// 	// os.Exit(0)
// 	return &playlist, nil
// }

// func handleInterrupt(ui *raw.TUI, playedList []string, completedPlaylist int) {
// 	raw.HideCursor()

// 	// handle CTRL C
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)

// 	go func() {
// 		for range c {
// 			ui.DisplayStats(playedList, completedPlaylist)
// 		}
// 	}()
// }

// // func UniqSong() (songNum int) {
// //
// // 	for {
// // 		songNum = rand.Intn(len(playlist))
// // 		if !UserSetting.RepeatSong {
// // 			if inPlaylist(songNum) {
// // 				continue
// // 			}
// // 		}
// // 		break
// // 	}
// //
// // 	return songNum
// // }

// // func inPlaylist(songNum int, playlist []string) bool {
// // 	for _, v := range playedList {
// // 		if playlist[songNum] == v {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// func getSong(i int, playlist *types.Playlist, setting types.Config) {
// 	var prevSong, curSong, nextSong int
// 	prevSong = i - 1
// 	curSong = i
// 	if len(playlist.List) == i+1 {
// 		nextSong = i
// 		if setting.Music.RepeatPlaylist {
// 			nextSong = 0
// 		}
// 	} else {
// 		nextSong = i + 1
// 	}
// 	playlist.CurrentSong = curSong
// 	playlist.NextSong = nextSong
// 	playlist.PrevSong = prevSong
// }

// func appendOnlyOriginal(list []string, val string) (originalList []string) {
// 	for _, v := range list {
// 		if v == val {
// 			return list
// 		}
// 	}
// 	originalList = append(list, val)
// 	return originalList
// }

// func addFiles(files []string) *types.Playlist {
// 	playlist = types.Playlist{}
// 	for _, file := range files {
// 		ext := filepath.Ext(file)
// 		if ext == ".mp3" {
// 			path, _ := filepath.Abs(filepath.Dir(file))
// 			playlist.List = append(playlist.List, types.Song{
// 				Name: filepath.Base(file),
// 				Path: path,
// 			})
// 		}
// 	}
// 	return &playlist
// }

// func excludeFiles(playlist *types.Playlist, exclude []string) *types.Playlist {
// 	for i, file := range playlist.List {
// 		full := filepath.Join(file.Path, file.Name)
// 		if containsS(exclude, full) {
// 			playlist.List = append(playlist.List[:i], playlist.List[i+1:]...)
// 		}
// 	}
// 	return playlist
// }

// func containsS(arr []string, str string) bool {
// 	for _, a := range arr {
// 		if a == str {
// 			return true
// 		}
// 	}
// 	return false
// }

// func contains(arr []types.Song, str string) bool {
// 	for _, a := range arr {
// 		full := filepath.Join(a.Path, a.Name)
// 		if full == str {
// 			return true
// 		}
// 	}
// 	return false
// }

// func includeFiles(playlist *types.Playlist, include []string) *types.Playlist {
// 	for _, file := range include {
// 		if !contains(playlist.List, file) {
// 			playlist.List = append(playlist.List, types.Song{
// 				Name: filepath.Base(file),
// 				Path: filepath.Dir(file),
// 			})
// 		}
// 	}
// 	return playlist
// }
