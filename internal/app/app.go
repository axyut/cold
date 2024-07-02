package app

import (
	"github.com/axyut/playgo/internal/booTea"
	"github.com/axyut/playgo/internal/rawtui"
	"github.com/axyut/playgo/internal/types"
)

func StartApp(config *types.Config) {
	// fmt.Println("Starting app with config: ", config)
	if config.Renderer == "raw" {
		rawtui.StartRawTui(config)
	} else {
		booTea.RunBubbleTUI(config)
	}
}
