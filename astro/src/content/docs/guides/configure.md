---
title: Configure
description: A guide to Install and Configure Playgo.
---

Configure the player according to your needs on Linux, Windows and Mac.

About settings...
A config file will be generated when you first run `playgo`. Depending on your operating system it can be found in one of the following locations:

-   macOS: `~/Library/Application\ Support/playgo/config.yml`
-   Linux: `~/.config/playgo/config.yml`
-   Windows: `C:\Users\me\AppData\Roaming\playgo\config.yml`

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

## Further reading

-   Read [reference](/playgo/reference/example) and Examples.
