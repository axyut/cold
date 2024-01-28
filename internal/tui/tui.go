package tui

import (
	"fmt"
	"os"
)

func Display(playlist []string, notifs []string, currentSong, prevSong, nextSong int, Shuffle, RepeatSong, RepeatPlaylist bool) {
	clear()
	HideCursor()
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
		prev := currentSong - i
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
	color.Cyan(fmt.Sprintf("⏯️%s", stripString(playlist[currentSong])))

	// next songs
	totalNextSongs := 6
	for j := 1; j <= totalNextSongs; j++ {
		next := currentSong + j
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
	fmt.Fprintf(screen, "Shuffle: %t", Shuffle)
	moveCursor(pos{3, intH + 1})
	fmt.Fprintf(screen, "Repeat Song: %t", RepeatSong)
	moveCursor((pos{3, intH + 2}))
	fmt.Fprintf(screen, "Repeat playlist: %t", RepeatPlaylist)

	// Now Playing
	moveCursor(pos{maxX / 2, 1})
	fmt.Fprintf(screen, "NOW PLAYING")
	moveCursor(pos{maxX / 2, 3})
	fmt.Fprintf(screen, "%s", stripString(playlist[currentSong]))
	moveCursor(pos{maxX / 2, 4})
	fmt.Fprintf(screen, "%d:00 -------------------- 3:14s")
	// song info
	// seek info

	// Notification
	moveCursor(pos{maxX / 2, int(float32(maxY)/1.25) - 1})
	fmt.Fprintln(screen, "NOTIFICATIONS")
	for i, v := range notifs {
		if i > 4 {
			break
		}
		moveCursor(pos{maxX / 2, int(float32(maxY)/1.25) + i})
		fmt.Fprintf(screen, " %s", stripString(v))
	}

	render()
}

func DisplayStats(playlist, playedList []string, completedPlaylist int) {
	clear()
	showCursor()

	moveCursor(pos{2, 2})
	fmt.Fprintf(screen, "Played         : %d song(s).", len(playedList)+(len(playlist)*completedPlaylist))
	moveCursor(pos{2, 3})
	fmt.Fprintf(screen, "Played list    : %d time(s).", completedPlaylist)
	moveCursor(pos{2, 4})
	fmt.Fprintf(screen, "Minutes played : 21 minute(s)")

	render()
	os.Exit(0)
}
