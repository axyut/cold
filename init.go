package main

import (
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"os"
)

// Prepare an Oto context (creating context for single time)
var op = &oto.NewContextOptions{
	SampleRate:   44100,
	ChannelCount: 2,
	Format:       oto.FormatSignedInt16LE,
}
var otoCtx, readyChan, otoErr = oto.NewContext(op)

var Playlist []string
var PlayedList []string
var Favorites []string
var Notifications []string
var Timer int
var CompletedPlaylist int

type pos [2]int

type Player struct {
	Music       *oto.Player
	UserSetting Setting
	File        *os.File
	Song        int
	MP3         mp3.Decoder
}

type Stats struct {
	MinutesPlayed int
	SongsPlayed   int
}

type Setting struct {
	Shuffle        bool
	RepeatSong     bool
	RepeatPlaylist bool
}

var UserSetting = Setting{
	true,
	false, // if Shuffle false it's no use for RepeatSong to be true
	true,
}

type Activelist struct {
	prevSong    int
	currentSong int
	nextSong    int
}

var Songs = Activelist{
	-1, 0, 1,
}
