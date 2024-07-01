package rawtui

import (
	"fmt"
	"log"
	"time"

	"github.com/axyut/playgo/internal/player"
	"github.com/axyut/playgo/internal/types"
	"github.com/axyut/playgo/pkg/raw"
)

func StartRawTui(cfg types.Config) {
	handleConfig(cfg)
	go listenForKey(cfg)
	ui := raw.NewUI(&playlist, &Notifications, &cfg)
	ui.HandleInterrupt(playedList, completedPlaylist)

top:
	for i := 0; ; i++ {
	repeatSameSong:
		// getSong(i, &playlist, cfg)
		playlist.CurrentSong = i
		player := player.NewPlayer(&playlist, &cfg)
		if player == nil {
			continue
		}

		go ui.Display()

		for player.Music.IsPlaying() {
			time.Sleep(time.Millisecond)
		}

		err := player.Music.Close()
		if err != nil {
			panic("player.Close failed: " + err.Error())
		}
		player.File.Close()

		playedList = appendOnlyOriginal(playedList, playlist.List[playlist.CurrentSong].Name)
		played := fmt.Sprintf("Played %s", playlist.List[playlist.CurrentSong].Name)
		notify(played)

		if len(playedList) >= len(playlist.List) {
			notify("Completed Playlist.")
			playedList = []string{}
			completedPlaylist++
			i = 0
			break
		}
		if cfg.Music.RepeatSong {
			notify("Playing Same Song Again.")
			goto repeatSameSong
		}
	}
	if cfg.Music.RepeatPlaylist {
		notify("Restarting Playlist.")
		goto top
	}
	ui.DisplayStats(playedList, completedPlaylist)
}

func handleConfig(config types.Config) {

	playlist, err := addFolder(config.General.StartDir)
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
		playlist = addFiles(config.Temp.PlayOnly)
	}

	if config.Music.Shuffle {
		shufflePlaylist(playlist)
	}
}
