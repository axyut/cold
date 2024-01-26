package tui

import (
	"bufio"
	"fmt"
	"os"

	"github.com/axyut/playgo"
	"golang.org/x/term"
)

type pos [2]int

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

type Player main.Player

func display(p *Player) {
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
