package list

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/axyut/playgo/internal/types"
)

type Playlist struct {
	List        []types.Song
	CurrentSong int
}

func NewPlaylist(config *types.Config) *Playlist {
	return handleConfig(*config)
}

func handleConfig(config types.Config) *Playlist {

	playlist, err := AddFolder(&Playlist{}, config.General.StartDir)
	if err != nil {
		log.Default().Println(err)
	}
	if len(config.Temp.Exclude) != 0 {
		playlist = excludeFiles(playlist, config.Temp.Exclude)
	}
	if len(config.Temp.Include) != 0 {
		playlist = includeFiles(playlist, config.Temp.Include)
	}
	if len(config.Temp.PlayOnly) != 0 {
		playlist = playlist.addFiles(config.Temp.PlayOnly)
	}

	if config.Music.Shuffle {
		playlist.Shuffle()
	}
	return playlist
}

func (playlist Playlist) Shuffle() {
	list := playlist.List
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}

func (playlist Playlist) Serialize(path string) error {
	// just doing addFolder for now which doesn't cover when individual files opened in command $playgo a.mp3 b.mp3
	if _, err := AddFolder(&playlist, path); err != nil {
		return err
	}
	return nil
}

func AddFolder(playlist *Playlist, path string) (*Playlist, error) {
	// get full path
	if path == "." {
		path, _ = filepath.Abs(path)
	}
	fileInfos, err := os.ReadDir(path)
	if err != nil {
		log.Println("Couldn't Read from ", path, " Directory!")
		// return &playlist, err
	}
	for _, file := range fileInfos {
		ext := filepath.Ext(file.Name())
		if ext == ".mp3" {
			// path, _ := filepath.Abs(filepath.Dir(file.Name()))
			song := types.Song{
				Name: file.Name(),
				Path: path,
			}

			playlist.List = append(playlist.List, song)
		}
	}
	home, _ := os.UserHomeDir()
	musicPath := filepath.Join(home + "/Music/")
	// fmt.Println("Music Path: ", musicPath)
	if len(playlist.List) == 0 && path == musicPath {
		fmt.Println("No Music Files Found in the given path: ", path)
		os.Exit(0)
	} else if len(playlist.List) == 0 {
		fmt.Println("No Music Files Found in the given path:", path, "Trying ~/Music/")
		_, err := AddFolder(playlist, musicPath)
		if err != nil {
			log.Println("Couldn't Read from ~/Music/ Directory!")
			return playlist, err
		}
	}
	// fmt.Println("Adding ", playlist.List)
	// os.Exit(0)
	return playlist, nil
}

func (playlist Playlist) addFiles(files []string) *Playlist {
	playlist = Playlist{}
	for _, file := range files {
		ext := filepath.Ext(file)
		if ext == ".mp3" {
			path, _ := filepath.Abs(filepath.Dir(file))
			playlist.List = append(playlist.List, types.Song{
				Name: filepath.Base(file),
				Path: path,
			})
		}
	}
	return &playlist
}

func excludeFiles(playlist *Playlist, exclude []string) *Playlist {
	for i, file := range playlist.List {
		full := filepath.Join(file.Path, file.Name)
		if containsS(exclude, full) {
			playlist.List = append(playlist.List[:i], playlist.List[i+1:]...)
		}
	}
	return playlist
}

func containsS(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func contains(arr []types.Song, str string) bool {
	for _, a := range arr {
		full := filepath.Join(a.Path, a.Name)
		if full == str {
			return true
		}
	}
	return false
}

func includeFiles(playlist *Playlist, include []string) *Playlist {
	for _, file := range include {
		if !contains(playlist.List, file) {
			playlist.List = append(playlist.List, types.Song{
				Name: filepath.Base(file),
				Path: filepath.Dir(file),
			})
		}
	}
	return playlist
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
