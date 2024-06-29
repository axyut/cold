package mp3decoder

import (
	"io"

	"github.com/hajimehoshi/go-mp3"
)

// will decode mp3 later without depending on this package
func Decode(file io.Reader) (*mp3.Decoder, error) {
	decoded, err := mp3.NewDecoder(file)
	return decoded, err
}
