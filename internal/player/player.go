package player

import (
	"log"
	"os"
	"path/filepath"

	"github.com/axyut/playgo/internal/list"
	"github.com/axyut/playgo/internal/types"
	mp3 "github.com/axyut/playgo/pkg/mp3Decoder"
	"github.com/ebitengine/oto/v3"
)

type Player struct {
	Music    *oto.Player
	Config   *types.Config
	File     *os.File
	Playlist *list.Playlist
}

// Prepare an Oto context (creating context for single time)
var op = &oto.NewContextOptions{
	SampleRate:   44100,
	ChannelCount: 2,
	Format:       oto.FormatSignedInt16LE,
}
var otoCtx, readyChan, otoErr = oto.NewContext(op)

// seek, next , prevous, pause, play, settings
func NewPlayer(playlist *list.Playlist, cfg *types.Config) *Player {
	// fmt.Println("NewPlayer", playlist)
	musicFile := playlist.List[playlist.CurrentSong]
	path := filepath.Join(musicFile.Path, musicFile.Name)
	// fmt.Println("Playing: ", path)
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
		// playlist = Remove(playlist.List, songNum)
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

	newPlayer := &Player{
		Music:    otoPlayer,
		Config:   cfg,
		File:     file,
		Playlist: playlist,
	}
	newPlayer.Music.Play()
	return newPlayer
}

func (player *Player) Pause() error {
	player.Music.Pause()
	return nil
}
