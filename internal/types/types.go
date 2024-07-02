package types

type Song struct {
	Name string
	Path string
}

// Config represents the main config for the application.
type Config struct {
	General  GeneralSettings `yaml:"general"`
	Player   PlayerSettings  `yaml:"player"`
	Music    MusicSettings   `yaml:"music"`
	Renderer string          `yaml:"renderer"`
	User     UserSetting     `yaml:"user"`
	Temp     TempSetting     `yaml:"temp"`
}

type GeneralSettings struct {
	StartDir      string `yaml:"start_dir"`
	ShowIcons     bool   `yaml:"show_icons"`
	ShowHidden    bool   `yaml:"show_hidden"`
	EnableLogging bool   `yaml:"enable_logging"`
}

type PlayerSettings struct {
	Shuffle        bool `yaml:"shuffle"`
	RepeatPlaylist bool `yaml:"repeat_playlist"`
}

type MusicSettings struct {
	RepeatSong     bool `yaml:"repeat_song"`
	RepeatPlaylist bool `yaml:"repeat_playlist"`
	Shuffle        bool `yaml:"shuffle"`
}

type UserSetting struct {
	UseDB bool `yaml:"use_db"`
}

type TempSetting struct {
	StartDir      string   `yaml:"start_dir"`
	ShowIcons     bool     `yaml:"show_icons"`
	ShowHidden    bool     `yaml:"show_hidden"`
	Exclude       []string `yaml:"exclude"`
	PlayOnly      []string `yaml:"playonly"`
	Include       []string `yaml:"include"`
	Renderer      string   `yaml:"renderer"`
	EnableLogging bool     `yaml:"enable_logging"`
}

type Activelist struct {
	PrevSong    int
	CurrentSong int
	NextSong    int
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
