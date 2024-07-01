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
-   Windows `go`
-   Debian `go libasound2-dev`
-   Mac
-   Arch
-   Fedora

## Installation

### Go

```bash
go install github.com/axyut/playgo@latest
```

### Curl

```bash
curl ...install.sh
```

## Development

-   `go run . .` from the source
-   `GOOS=linux GOARCH=amd64 go build -o ~/go/bin/ .` build to your bin path.
-   `GOOS=linux GOARCH=amd64 go build . && sudo cp playgo /usr/local/bin/` build for the system.
    -   `GOOS=darwin GOARCH=amd64` Mac
    -   `GOOS=windows GOARCH=amd64` Windows

## Features

-   [x] Plays Music.
-   [x] Light weight with minimum dependencies.
-   [ ] Controls.
-   [ ] Stream, Radio
-   [ ] Visualization.

## Themes

-   [ ] UI with [bubbletea](https://github.com/charmbracelet/bubbletea)
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

```plaintext
## flags
  play files                  - $playgo <file.mp3> <file2.mp3>
  play all music in folder    - $playgo / $playgo . / $playgo ~/Music/path
  help                        - $playgo -h
  test condition/health       - $playgo -t
## while playing
  q - quit player
  p - Play/Pause

  h - play previous song
  j - seek backward 10s
  k - seek forward 10s
  l - play next song

  w - Increase Volume by 5%
  a -
  s - Decrease Volume by 5%
  d -

  e - Toogle Repeat Playlist On/Off
  r - Toogle Repeat Song On/Off
  t - Toogle Shuffle On/Off
```

## Configuration

About settings...
A config file will be generated when you first run `playgo`. Depending on your operating system it can be found in one of the following locations:

-   macOS: ~/Library/Application\ Support/playgo/config.yml
-   Linux: ~/.config/playgo/config.yml
-   Windows: C:\Users\me\AppData\Roaming\playgo\config.yml

It will include the following default settings:

```yml
setting:
    general:
        show_icons: true
    player:
        shuffle: true
        repeat_playlist: true
    music:
        repeat_song: false
theme:
    raw: true
```

## Uninstalling

If installed with `go install` remove bin file from your go bin path. If not setup, by default it is in `~/go/bin`, If installed with curl ...

<!-- imp
[project structure](https://github.com/golang-standards/project-layout/blob/master/test/README.md)
[github actions trigger](https://github.com/termkit/gama?tab=readme-ov-file)
[commands parser](https://github.com/alecthomas/kingpin)
[go releaser](https://github.com/wangyoucao577/go-release-action)
setup go build command with test mp3 files

git tag v0.1.1
git push origin v0.1.1
GOPROXY=proxy.golang.org go list -m github.com/axyut/playgo@v0.1.1
-->
