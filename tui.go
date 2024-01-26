package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

var screen = bufio.NewWriter(os.Stdout)

func hideCursor() {
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
func (p *Player) display() {
	clear()
	hideCursor()
	// border()
	seprator()
	maxX, maxY := termSize()

	// Playlist
	// TODO: will make playlist scrollable with cursor later on
	// will later implement this to show all playlist when certain key pressed.
	moveCursor(pos{3, 1})
	fmt.Fprintf(screen, "PLAYLIST (%d songs)", len(playlist))

	// 1. iterate playedList? start currentSong from top and move down as playedList increases upto ~5 prev songs
	// 2. iterate playlist? start at constant row and provide what would be last ~5 songs then gradually add playedList
	// prev songs, idk which option better, now implementing 2.
	totalPrevSongs := 3
	for i := 1; i <= totalPrevSongs; i++ {
		prev := songs.currentSong - i
		if prev <= -1 {
			for prev <= -1 {
				prev = prev + len(playlist)
			}
		}
		moveCursor(pos{2, (maxY / 4) - i})
		color.Magenta(fmt.Sprintf("%s", stripString(playlist[prev])))
	}

	// currently Playing
	moveCursor(pos{2, maxY / 4})
	color.Reversed()
	color.Cyan(fmt.Sprintf("⏯️%s", stripString(playlist[songs.currentSong])))

	// next songs
	totalNextSongs := 6
	for j := 1; j <= totalNextSongs; j++ {
		next := songs.currentSong + j
		if next >= len(playlist) {
			for next >= len(playlist) {
				next = next - len(playlist)
			}
		}
		moveCursor(pos{2, (maxY / 4) + j})
		color.Blue(fmt.Sprintf("%s", stripString(playlist[next])))
	}

	// Settings
	intH := int(float32(maxY) / 1.25)
	moveCursor(pos{2, intH - 1})
	fmt.Fprintf(screen, "SETTINGS")
	moveCursor(pos{3, intH})
	fmt.Fprintf(screen, "Shuffle: %t", UserSetting.Shuffle)
	moveCursor(pos{3, intH + 1})
	fmt.Fprintf(screen, "Repeat Song: %t", UserSetting.RepeatSong)
	moveCursor((pos{3, intH + 2}))
	fmt.Fprintf(screen, "Repeat playlist: %t", UserSetting.RepeatPlaylist)

	// Now Playing
	moveCursor(pos{maxX / 2, 1})
	fmt.Fprintf(screen, "NOW PLAYING")
	moveCursor(pos{maxX / 2, 3})
	fmt.Fprintf(screen, "%s", stripString(p.File.Name()))
	moveCursor(pos{maxX / 2, 4})
	fmt.Fprintf(screen, "%d:00 -------------------- 3:14s", timer)
	// song info
	// seek info

	// Notification
	moveCursor(pos{maxX / 2, int(float32(maxY)/1.25) - 1})
	fmt.Fprintln(screen, "NOTIFICATIONS")
	for i, v := range notifications {
		if i > 4 {
			break
		}
		moveCursor(pos{maxX / 2, int(float32(maxY)/1.25) + i})
		fmt.Fprintf(screen, " %s", stripString(v))
	}

	render()
}

func displayStats() {
	clear()
	showCursor()

	// moveCursor(pos{1, 1})
	// fmt.Fprintf(screen, "Stats :")
	moveCursor(pos{2, 2})
	fmt.Fprintf(screen, "Played         : %d song(s).", len(playedList)+(len(playlist)*completedPlaylist))
	moveCursor(pos{2, 3})
	fmt.Fprintf(screen, "Played list    : %d time(s).", completedPlaylist)
	moveCursor(pos{2, 4})
	fmt.Fprintf(screen, "Minutes played : 21 minute(s)")

	render()
	os.Exit(0)
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
