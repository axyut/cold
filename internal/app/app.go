package app

import (
	"fmt"
	"log"
	"os"
	"time"

	c "github.com/axyut/playgo/internal/config"
	"github.com/axyut/playgo/internal/tui"
	mp3 "github.com/axyut/playgo/pkg/mp3Decoder"
)

func StartPlaygo(cfg c.Config) {
	handleConfig(cfg)
	go listenForKey(cfg)
	ui := tui.NewUI(&playlist, &Notifications, &songs, &cfg)
	handleInterrupt(ui, playedList, completedPlaylist)

top:
	for i := 0; ; i++ {
	repeatSameSong:
		songs := getSong(i, &playlist, cfg)

		player := play(songs.CurrentSong, &cfg)
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

		playedList = appendOnlyOriginal(playedList, playlist[songs.CurrentSong])
		played := fmt.Sprintf("Played %s", playlist[songs.CurrentSong])
		notify(played)

		if len(playedList) == len(playlist) {
			playedList = []string{}
			completedPlaylist++
			break
		}
		if cfg.Music.RepeatSong {
			goto repeatSameSong
		}
	}
	if cfg.Music.RepeatPlaylist {
		notify("Restarting Playlist.")
		goto top
	}
	ui.DisplayStats(playedList, completedPlaylist)
}

// seek, next , prevous, pause, play, settings
func play(songNum int, cfg *c.Config) *Player {
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
		cfg,
		file,
		songNum,
	}
	newPlayer.Music.Play()
	return &newPlayer
}

func handleConfig(config c.Config) {
	err := addFolder(config.General.StartDir, &playlist)
	if err != nil {
		log.Default().Println(err)
	}
	// check if it's files or a folder
	if config.Music.Shuffle {
		shufflePlaylist(&playlist)
	}
}
