package app

import (
	"github.com/axyut/cold/internal/booTea"
	"github.com/axyut/cold/internal/rawtui"
	"github.com/axyut/cold/internal/types"
)

func StartApp(config *types.Config) {
	// fmt.Println("Starting app with config: ", config)
	if config.Renderer == "raw" {
		rawtui.StartRawTui(config)
	} else {
		booTea.RunBubbleTUI(config)
	}
}
