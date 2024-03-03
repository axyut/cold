package tui

import (
	"fmt"
	"os"
	"time"

	g "github.com/axyut/playgo/internal/global"
)

var maxX, maxY = termSize()

func Display(rplaylist *[]string, rnotifs *[]string, songs *g.Activelist, UserSetting *g.Setting) {
	for {
		playlist := *rplaylist
		notifs := *rnotifs
		currentSong := *&songs.CurrentSong
		Shuffle := *&UserSetting.Shuffle
		RepeatSong := *&UserSetting.RepeatSong
		RepeatPlaylist := *&UserSetting.RepeatPlaylist

		clear()
		HideCursor()
		// border()
		seprator()

		currentlyPlaying(playlist, currentSong)
		displayNextSongs(playlist, currentSong)
		displaySettings(Shuffle, RepeatSong, RepeatPlaylist)
		displayNowPlaying(playlist, currentSong)
		displayNotifications(notifs)

		render()
		time.Sleep(time.Millisecond * 500)
	}
}

func currentlyPlaying(playlist []string, currentSong int) {
	moveCursor(pos{2, maxY / 4})
	color.Reversed()
	color.Cyan(fmt.Sprintf("⏯️%s", stripString(playlist[currentSong])))
}

func displayNextSongs(playlist []string, currentSong int) {
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

}

func displaySettings(Shuffle, RepeatSong, RepeatPlaylist bool) {
	intH := int(float32(maxY) / 1.25)
	moveCursor(pos{2, intH - 1})
	fmt.Fprintf(screen, "SETTINGS")
	moveCursor(pos{3, intH})
	fmt.Fprintf(screen, "[t] Shuffle: %t", Shuffle)
	moveCursor(pos{3, intH + 1})
	fmt.Fprintf(screen, "[r] Repeat Song: %t", RepeatSong)
	moveCursor((pos{3, intH + 2}))
	fmt.Fprintf(screen, "[e] Repeat playlist: %t", RepeatPlaylist)
}

func displayNowPlaying(playlist []string, currentSong int) {
	moveCursor(pos{maxX / 2, 1})
	fmt.Fprintf(screen, "NOW PLAYING")
	moveCursor(pos{maxX / 2, 3})
	fmt.Fprintf(screen, "%s", stripString(playlist[currentSong]))
	moveCursor(pos{maxX / 2, 4})
	fmt.Fprintf(screen, "0:00 -------------------- 3:14s")
}

func displayNotifications(notifs []string) {
	moveCursor(pos{maxX / 2, int(float32(maxY)/1.25) - 1})
	fmt.Fprintln(screen, "NOTIFICATIONS")
	for i, v := range notifs {
		if i > 4 {
			break
		}
		moveCursor(pos{maxX / 2, int(float32(maxY)/1.25) + i})
		fmt.Fprintf(screen, " %s", stripString(v))
	}
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
