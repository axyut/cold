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
func (p *Player) display() {
	// border()
	// for loop ?  and Sleep
	// var timer int
	// for {
	// timer++
	maxX, maxY := termSize()
	moveCursor(pos{3, 1})
	fmt.Fprintf(screen, "PLAYLIST (%d songs)", len(playlist))
	for i, v := range playlist {
		stripped := stripString(v)
		if i > maxY/(3) {
			moveCursor(pos{2, i + 3})
			fmt.Fprintf(screen, "  %d more Songs...", len(playlist)-i)
			moveCursor(pos{2, i + 4})
			fmt.Fprintf(screen, "%d. %s", len(playlist), stripString(playlist[len(playlist)-1]))
			break
		}
		moveCursor(pos{2, i + 3})
		fmt.Fprintf(screen, "%d. %s", i+1, stripped)
	}

	moveCursor(pos{2, int(float32(maxY)/1.25) - 1})
	fmt.Fprintf(screen, "SETTINGS")

	moveCursor(pos{maxX / 2, 1})
	fmt.Fprintf(screen, "NOW PLAYING")
	moveCursor(pos{maxX / 2, 3})
	fmt.Fprintf(screen, "%s", stripString(p.File.Name()))
	moveCursor(pos{maxX / 2, 4})
	fmt.Fprintf(screen, "Length: %d ---------------- %d seconds", timer, p.MP3.Length()/60/60/60)

	moveCursor(pos{maxX / 2, maxY / 4})
	fmt.Fprintf(screen, "NEXT SONG")
	moveCursor(pos{maxX / 2, (maxY / 4) + 1})
	fmt.Fprintf(screen, "------")

	// Notification
	moveCursor(pos{maxX / 2, int(float32(maxY)/1.25) - 1})
	fmt.Fprintln(screen, "NOTIFICATIONS")
	for i, v := range notifications {
		if i > 4 {
			break
		}
		moveCursor(pos{maxX / 2, int(float32(maxY)/1.25) + i})
		fmt.Fprintf(screen, " %d. %s", i+1, v)

	}

	render()
	// time.Sleep(time.Second)
	// break
	// }
}

func displayStats() {
	clear()
	showCursor()
	moveCursor(pos{1, 1})
	fmt.Fprintf(screen, "Statistics")
	render()
	os.Exit(0)
}
