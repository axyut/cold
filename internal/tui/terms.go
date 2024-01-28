package tui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type pos [2]int

type Setting struct {
	Shuffle        bool
	RepeatSong     bool
	RepeatPlaylist bool
}

type Activelist struct {
	prevSong    int
	currentSong int
	nextSong    int
}

var screen = bufio.NewWriter(os.Stdout)

func HideCursor() {
	fmt.Fprint(screen, "\033[?25l")
}

func showCursor() {
	fmt.Fprint(screen, "\033[?25h")
}
func moveCursor(pos [2]int) {
	fmt.Fprintf(screen, "\033[%d;%dH", pos[1], pos[0])
}

func clear() {
	fmt.Fprint(screen, "\033[2J")
}

func draw(str string) {
	fmt.Fprint(screen, str)
}

// write all data in buffer to terminal
func render() {
	screen.Flush()
}

func termSize() (width int, height int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return width, height
}

func border() {
	maxX, maxY := termSize()
	// ---- top ---- bottom
	for i := 1; i <= maxX; i++ {
		fmt.Fprintf(screen, "\033[1;%dH", i)
		fmt.Fprintf(screen, "_")
		fmt.Fprintf(screen, "\033[%d;%dH", maxY, i)
		fmt.Fprintf(screen, "_")
	}
	// | left | right |
	for i := 1; i <= maxY; i++ {
		fmt.Fprintf(screen, "\033[%d;1H", i)
		fmt.Fprintf(screen, "|")
		fmt.Fprintf(screen, "\033[%d;%dH", i, maxX)
		fmt.Fprintf(screen, "|")
	}
}

func seprator() {
	maxX, maxY := termSize()
	for i := 1; i <= maxY; i++ {
		fmt.Fprintf(screen, "\033[%d;%dH", i, (maxX/2)-2)
		fmt.Fprintf(screen, "|")
	}
}

func resetColor() {
	fmt.Fprintf(screen, "\u001b[0m")
}

func stripString(str string) string {
	maxX, _ := termSize()
	strip := maxX / 3
	len := len(str)
	if isMp3 := strings.ContainsAny(str, "mp3"); isMp3 {
		str = str[:len-3]
		len = len - 3
	}
	if len > strip {
		return str[:strip] + "..."
	}
	return str
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
var color = Render{
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
