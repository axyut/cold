package global

type Activelist struct {
	PrevSong    int
	CurrentSong int
	NextSong    int
}

type Setting struct {
	Shuffle        bool
	RepeatSong     bool
	RepeatPlaylist bool
}

type Stats struct {
	MinutesPlayed int
	SongsPlayed   int
}

type Flag struct {
	Help string
	Test string
}

const Usage = `Usage
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
t - Toogle Shuffle On/Off`
