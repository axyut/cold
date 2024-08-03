---
title: Configure
description: A guide to Install and Configure cold.
---

Configure the player according to your needs on Linux, Windows and Mac.

About settings...
A config file will be generated when you first run `cold`. Depending on your operating system it can be found in one of the following locations:

-   macOS: `~/Library/Application\ Support/cold/config.yml`
-   Linux: `~/.config/cold/config.yml`
-   Windows: `C:\Users\me\AppData\Roaming\cold\config.yml`

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

-   Read [reference](/cold/reference/example) and Examples.
