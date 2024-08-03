package config

import (
	"os"
	"path/filepath"

	"github.com/axyut/cold/internal/types"
	"gopkg.in/yaml.v3"
)

const AppDir = "cold"
const ConfigFileName = "config.yml"

// configError represents an error that occurred while parsing the config file.
type configError struct {
	configDir string
	parser    ConfigParser
	err       error
}

// ConfigParser is the parser for the config file.
type ConfigParser struct{}

// initParser initializes the parser.
func initParser() ConfigParser {
	return ConfigParser{}
}

// ParseConfig parses the config file and returns the config.
func Parse(temp *types.TempSetting) (types.Config, error) {
	var config types.Config
	var err error

	parser := initParser()

	configFilePath, err := parser.getConfigFileOrCreateIfMissing()
	if err != nil {
		return config, parsingError{err: err}
	}

	config, err = parser.readConfigFile(*configFilePath)
	if err != nil {
		return config, parsingError{err: err}
	}
	config.Temp = *temp
	// log.Println("config: ", config)
	if temp.StartDir != "" {
		config.General.StartDir = temp.StartDir
	}
	if temp.Renderer == "raw" || temp.Renderer == "tea" {
		config.Renderer = temp.Renderer
	}
	if !temp.ShowIcons {
		config.General.ShowIcons = temp.ShowIcons
	}
	if temp.EnableLogging {
		config.General.EnableLogging = temp.EnableLogging
	}
	if temp.ShowHidden {
		config.General.ShowHidden = temp.ShowHidden
	}

	return config, nil
}

// getConfigFileOrCreateIfMissing returns the config file path or creates the config file if it doesn't exist.
func (parser ConfigParser) getConfigFileOrCreateIfMissing() (*string, error) {
	var err error
	configDir := os.Getenv("XDG_CONFIG_HOME")

	if configDir == "" {
		configDir, err = os.UserConfigDir()
		if err != nil {
			return nil, configError{parser: parser, configDir: configDir, err: err}
		}
	}

	prsConfigDir := filepath.Join(configDir, AppDir)
	err = os.MkdirAll(prsConfigDir, os.ModePerm)
	if err != nil {
		return nil, configError{parser: parser, configDir: configDir, err: err}
	}

	configFilePath := filepath.Join(prsConfigDir, ConfigFileName)
	err = parser.createConfigFileIfMissing(configFilePath)
	if err != nil {
		return nil, configError{parser: parser, configDir: configDir, err: err}
	}

	return &configFilePath, nil
}

// readConfigFile reads the config file and returns the config.
func (parser ConfigParser) readConfigFile(path string) (types.Config, error) {
	config := parser.getDefaultConfig()
	data, err := os.ReadFile(path)
	if err != nil {
		return config, configError{parser: parser, configDir: path, err: err}
	}

	err = yaml.Unmarshal((data), &config)
	parser.sanitizeConfig(&config)
	return config, err
}

func (parser ConfigParser) sanitizeConfig(config *types.Config) {
	if config.General.StartDir == "" {
		config.General.StartDir = "~/Music/"
	}
	if config.Renderer != "tea" && config.Renderer != "raw" {
		config.Renderer = "tea"
	}
	// type check gaera kaam xaina kina ki yaml.Unmarshal le type check garxa
	// if fmt.Sprintf("%T", config.General.showIcons) != "bool" {
	// 	config.General.showIcons = true
	// }
	// if fmt.Sprintf("%T", config.General.EnableLogging) != "bool" {
	// 	config.General.EnableLogging = false
	// }
	// if fmt.Sprintf("%T", config.Player.Shuffle) != "bool" {
	// 	config.Player.Shuffle = true
	// }
	// if fmt.Sprintf("%T", config.Player.RepeatPlaylist) != "bool" {
	// 	config.Player.RepeatPlaylist = true
	// }
	// if fmt.Sprintf("%T", config.Music.RepeatSong) != "bool" {
	// 	config.Music.RepeatSong = false
	// }

}

// getDefaultConfig returns the default config for the application.
func (parser ConfigParser) getDefaultConfig() types.Config {
	return types.Config{

		General: types.GeneralSettings{
			ShowIcons:     true,
			ShowHidden:    false,
			StartDir:      "~/Music/",
			EnableLogging: false,
		},
		Player: types.PlayerSettings{
			Shuffle:        true,
			RepeatPlaylist: true,
		},
		Music: types.MusicSettings{
			RepeatSong:     false,
			RepeatPlaylist: true,
			Shuffle:        true,
		},
		Renderer: "raw",
		User: types.UserSetting{
			UseDB: false,
		},
	}
}
