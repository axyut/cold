package rawtui

import (
	"fmt"
	"time"

	"github.com/axyut/playgo/internal/list"
	"github.com/axyut/playgo/internal/player"
	"github.com/axyut/playgo/internal/types"
	"github.com/axyut/playgo/pkg/raw"
)

func StartRawTui(cfg *types.Config) {
	playlist := list.NewPlaylist(cfg)
	ui := raw.NewUI(playlist, &Notifications, cfg)
	ui.HandleInterrupt(playedList, completedPlaylist)
top:
	for i := 0; ; i++ {
	repeatSameSong:
		// getSong(i, &playlist, cfg)
		playlist.CurrentSong = i
		player := player.NewPlayer(playlist, cfg)
		if player == nil {
			continue
		}

		ui.ListenForKey(keysWithAction(playlist, cfg, player))
		ui.Display()

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

func keysWithAction(playlist *list.Playlist, setting *types.Config, player *player.Player) []raw.ListenKeyAndAction {

	return []raw.ListenKeyAndAction{
		{
			Key: 'a',
			Action: func() {
				notify("<< Previous")
			},
		},
		{
			Key: 'q',
			Action: func() {
				notify("<< 10s ")
			},
		},
		{
			Key: 'e',
			Action: func() {
				player.Music.Seek(10, 0)
				notify(">> 10s")
			},
		},
		{
			Key: 'd',
			Action: func() {
				player.Music.Close()
				notify(">> Next")
			},
		},
		{
			Key: 'p',
			Action: func() {
				if player.Music.IsPlaying() {
					player.Music.Pause()
					notify("|> Paused")
				} else {
					player.Music.Play()
					notify("|| Playing")
				}
			},
		},
		{
			Key: 'w',
			Action: func() {
				notify("++ VOL")
			},
		},
		{
			Key: 's',
			Action: func() {
				notify("-- VOL")
			},
		},
		{
			Key: 'y',
			Action: func() {
				toogleSetting('y', playlist, setting)
			},
		},
		{
			Key: 't',
			Action: func() {
				toogleSetting('t', playlist, setting)
			},
		},
		{
			Key: 'r',
			Action: func() {
				toogleSetting('r', playlist, setting)
			},
		},
	}

}

func toogleSetting(str rune, list *list.Playlist, setting *types.Config) {
	suf, repS, repP := setting.Music.Shuffle, setting.Music.RepeatSong, setting.Music.RepeatPlaylist
	switch str {
	case 't':
		repP = !repP
	case 'r':
		repS = !repS
	case 'y':
		// serialize
		suf = !suf
		if suf {
			list.Shuffle()
			notify("Shuffle On.")
		} else {
			err := list.Serialize(setting.General.StartDir)
			if err != nil {
				notify("Error in serializing.")
			} else {
				notify("Shuffle Off.")
			}
		}
	}
	setting.Music.Shuffle = suf
	setting.Music.RepeatSong = repS
	setting.Music.RepeatPlaylist = repP
}
