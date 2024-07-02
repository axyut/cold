package rawtui

var Notifications []string

var playedList []string
var completedPlaylist int

func notify(str string) {
	Notifications = append([]string{str}, Notifications...)
}

func appendOnlyOriginal(list []string, val string) (originalList []string) {
	for _, v := range list {
		if v == val {
			return list
		}
	}
	originalList = append(list, val)
	return originalList
}
