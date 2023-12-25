package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

// Prepare an Oto context (creating context for single time)
var op = &oto.NewContextOptions{
	SampleRate:   44100,
	ChannelCount: 2,
	Format:       oto.FormatSignedInt16LE,
}
var otoCtx, readyChan, err = oto.NewContext(op)

func main() {
	// TODO:while playing display visual
	// TODO:custom mp3 decoder and player
	if len(os.Args) == 1 {
		panic("Open a .mp3 file.")
	}
	// check if it's a single a folder
	if os.Args[1] == "." {
		playFolder()
	} else {
		mp3File := os.Args[1]
		play(mp3File)
	}
}

func play(mp3File string) {
	file, err := os.Open(mp3File)
	if err != nil {
		panic("opening " + mp3File + " failed: " + err.Error())
	}

	fmt.Printf("-------- NOW NPLAYING %s --------\n", mp3File)
	// Decode file. This process is done as the file plays so it won't load whole thing into memory
	decodedMp3, err := mp3.NewDecoder(file)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}
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

func playFolder() {
	fmt.Printf("All Music from Current folder added to Playlist.\n")
	fileInfos, err := os.ReadDir(".")
	if err != nil {
		log.Fatal("Couldn't Read from Current Directory!\n")
	}

	var mp3Files []string
	for _, file := range fileInfos {
		ext := filepath.Ext(file.Name())
		if ext == ".mp3" {
			mp3Files = append(mp3Files, file.Name())
		}
	}
	for i, file := range mp3Files {
		fmt.Printf("%d. %s  ", i+1, file)
	}
	fmt.Printf("\n")
	for range mp3Files {
		// Shuffle Mode
		songNum := rand.Intn(len(mp3Files))
		play(mp3Files[songNum])
	}
}
