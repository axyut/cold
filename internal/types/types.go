package types

type Playlist struct {
	List        []Song
	CurrentSong int
}
type Song struct {
	Name string
	Path string
}
