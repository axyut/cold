<p align="center">
  <img src="./assets/logo.png" height="70" width="90" />
  <p align="center">
    CLI Music Player
  </p>
  <p align="center">
    <a href="https://github.com/axyut/playgo/releases"><img src="https://img.shields.io/github/v/release/axyut/playgo" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/axyut/playgo#section-readme"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
  </p>
</p>

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
  <img src="./assets/player.png" alt="default screenshot">
</p>

## About Playgo

Play Music in terminal. Written in Go. Relatively feature-packed.

### Requirements

-   [Go](https://golang.org/)
-   Windows
-   Debian `libasound2-dev`
-   Mac
-   Arch
-   Fedora `alsa-lib-devel`

## Installation

### Go

```bash
go install github.com/axyut/playgo@latest
```

### Curl

```bash
curl ...install.sh
```

## Build From Source

-   `go` version 1.22
-   `git clone https://github.com/axyut/playgo.git && cd playgo`
-   `make install` for debian.
-   `make build` or `make build_arm` for binary at `bin/`.
-   `cp bin/[filename] ~/.local/bin/playgo `
-   `playgo`

## Features

-   [x] Plays Music.
-   [x] Light weight with minimum dependencies.
-   [ ] Controls.
-   [ ] Stream, Radio
-   [ ] Visualization.

## Themes

-   [ ] UI with [bubbletea](https://github.com/charmbracelet/bubbletea)
-   [ ] Tview(https://github.com/rivo/tview)
-   [x] Raw UI Build

## Audio Engine

-   [x] [oto](https://github.com/ebitengine/oto/v3)
-   [x] [go-mp3](https://github.com/hajimehoshi/go-mp3)
-   [ ] [beep](https://github.com/faiface/beep)
-   [ ] [portaudio](https://github.com/gordonklaus/portaudio)

### Built With

-   [Go](https://golang.org/)
-   [Mp3 Engine (go-mp3, oto)](https://github.com/ebitengine/oto/v3)
-   [Cobra](https://github.com/spf13/cobra)
-   UI with [bubbletea](https://github.com/charmbracelet/bubbletea)
-   [Go FM](https://github.com/ssnat/GoFM)

## Usage

`playgo -h`

## Configuration

About settings...
A config file will be generated when you first run `playgo`. Depending on your operating system it can be found in one of the following locations:

-   macOS: ~/Library/Application\ Support/playgo/config.yml
-   Linux: ~/.config/playgo/config.yml
-   Windows: C:\Users\me\AppData\Roaming\playgo\config.yml

It will include the following default settings:

```yml
general:
    start_dir: /home/user/Music/score/
    enable_logging: true
player:
    shuffle: true
    repeat_playlist: true
music:
    repeat_song: false
    repeat_playlist: true
    shuffle: true
renderer: tea # tea, tview, raw
user:
    use_db: false
temp:
    start_dir: ""
    show_icons: false
    show_hidden: false
    exclude: []
    playonly: []
    include: []
    renderer: ""
    enable_logging: false
```

## Uninstalling

If installed with `go install` remove bin file from your go bin path. If not setup, by default it is in `~/go/bin`, If installed with curl ...

### IMP Links

[similar in go, gomu](https://github.com/issadarkthing/gomu)
[simlar in go, grump](https://github.com/dhulihan/grump)
[similar in rust](https://github.com/tramhao/termusic)
[project structure](https://github.com/golang-standards/project-layout/blob/master/test/README.md)
[github actions trigger](https://github.com/termkit/gama?tab=readme-ov-file)
[commands parser](https://github.com/alecthomas/kingpin)
[go releaser](https://github.com/wangyoucao577/go-release-action)
[radio example](https://github.com/vergonha/garden-tui)
[tui vis example](github.com/maaslalani/confetty)
[gstreamer](https://github.com/go-gst/go-gst)
[oto wrapper](https://github.com/gopxl/beep)

<!--
git tag v0.1.1
git push origin v0.1.1

only for first time
GOPROXY=proxy.golang.org go list -m github.com/axyut/playgo@v0.1.1
-->
