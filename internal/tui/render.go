package tui

import "fmt"

const Usage = `Usage
## flags
play files                  - $playgo <file.mp3> <file2.mp3>
play all music in folder    - $playgo .
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

var Flags = Flag{
	"h",
	"t",
}

type Render struct {
	Black     func(...string)
	Red       func(...string)
	Green     func(...string)
	Yellow    func(...string)
	Blue      func(...string)
	Magenta   func(...string)
	Cyan      func(...string)
	White     func(...string)
	BgBlack   func(...string)
	BgRed     func(...string)
	BgGreen   func(...string)
	BgYellow  func(...string)
	BgBlue    func(...string)
	BgMagenta func(...string)
	BgCyan    func(...string)
	BgWhite   func(...string)
	Bold      func(...string)
	Underline func(...string)
	Reversed  func(...string)
	Reset     func()
}

// works for now, not practical approach for adding new rendering methods
// color struct itoa or sth, map with string and display with one function without this shit.
var Color = Render{
	Bold: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[1m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[1m")
		}
	},
	Underline: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[4m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[4m")
		}
	},
	Reversed: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[7m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[7m")
		}
	},
	BgBlack: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[40m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[40m")
		}
	},
	BgRed: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[41m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[41m")
		}
	},
	BgGreen: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[42m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[42m")
		}
	},
	BgYellow: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[43m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[43m")
		}
	},
	BgBlue: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[44m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[44m")
		}
	},
	BgMagenta: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[45m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[45m")
		}
	},
	BgCyan: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[46m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[46m")
		}
	},
	BgWhite: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[47m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[47m")
		}
	},
	Black: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[30m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[30m")
		}
	},
	Red: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[31m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[31m")
		}
	},
	Green: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[32m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[32m")
		}
	},
	Yellow: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[33m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[33m")
		}
	},
	Blue: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[34m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[34m")
		}
	},
	Magenta: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[35m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[35m")
		}
	},
	Cyan: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[36m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[36m")
		}
	},
	White: func(text ...string) {
		if len(text) > 0 {
			for i := 0; i < len(text); i++ {
				fmt.Fprintf(screen, "\u001b[37m%s", text[i])
			}
			fmt.Fprintf(screen, "\u001b[0m")
		} else {
			fmt.Fprintf(screen, "\u001b[37m")
		}
	},
	Reset: func() {
		fmt.Fprintf(screen, "\u001b[0m")
	},
}
