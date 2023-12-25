package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

func main() {
	// TODO: open music file  and play
	// TODO: while playing display visual
	if len(os.Args) == 1 {
		panic("Open a .mp3 file.")
	}
	// check if it's a single a folder
	if os.Args[1] == "." {
		playFolder()
	}
	play()
	mp3File := os.Args[1]

	// Open the file for reading. Do NOT close before you finish playing!
	file, err := os.Open(mp3File)
	if err != nil {
		panic("opening " + mp3File + " failed: " + err.Error())
	}

	fmt.Printf("GoPlay.mp3 NOW NPLAYING %s", mp3File)
	// Decode file. This process is done as the file plays so it won't load whole thing into memory
	decodedMp3, err := mp3.NewDecoder(file)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	// Prepare an Oto context (this will use your default audio device) that will play
	op := &oto.NewContextOptions{}
	op.SampleRate = 44100 // usually 44100 or 48000
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE

	// Remember that you should **not** create more than one context
	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan
	player := otoCtx.NewPlayer(decodedMp3)
	player.Play()

	// We can wait for the sound to finish playing using something like this
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
	err = player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
	file.Close()
}

func play() {

}

func playFolder() {

}
