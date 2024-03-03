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
x - quit player
p - Play/Pause

q - seek backward 10s
e - seek forward 10s

w - Increase Volume by 5%
a - play previous song
s - Decrease Volume by 5%
d - play next song

r - Toogle Repeat Song On/Off
t - Toogle Repeat Playlist On/Off
y - Toogle Shuffle On/Off`
