package main

import (
	"os"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

// Prepare an Oto context (creating context for single time)
var op = &oto.NewContextOptions{
	SampleRate:   44100,
	ChannelCount: 2,
	Format:       oto.FormatSignedInt16LE,
}
var otoCtx, readyChan, otoErr = oto.NewContext(op)
var playlist []string
var playedList []string
var favorites []string
var notifications []string
var timer int
var completedPlaylist int

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

var songs = Activelist{
	-1, 0, 1,
}

const usage = `Usage
## flags
play files                  - $playgo <file.mp3> <file2.mp3>
play all music in folder    - $playgo / $playgo . / $playgo ~/Music/path
help                        - $playgo -h
test condition/health       - $playgo -t
## while playing
q - quit player
p - Play/Pause

h - play previous song
j - seek backward 10s
k - seek forward 10s
l - play next song

w - Increase Volume by 5%
a -
s - Decrease Volume by 5%
d -

e - Toogle Repeat Playlist On/Off
r - Toogle Repeat Song On/Off
t - Toogle Shuffle On/Off
`

type Flag struct {
	Help string
	Test string
}

var flags = Flag{
	"h",
	"t",
}
