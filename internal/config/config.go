package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const AppDir = "playgo"
const ConfigFileName = "config.yml"

// Config represents the main config for the application.
type Config struct {
	General  generalSettings `yaml:"general"`
	Player   playerSettings  `yaml:"player"`
	Music    musicSettings   `yaml:"music"`
	Renderer string          `yaml:"renderer"`
	User     userSetting     `yaml:"user"`
	Temp     TempSetting     `yaml:"temp"`
}

type generalSettings struct {
	showIcons     bool   `yaml:"show_icons"`
	StartDir      string `yaml:"start_dir"`
	EnableLogging bool   `yaml:"enable_logging"`
}

type playerSettings struct {
	Shuffle        bool `yaml:"shuffle"`
	RepeatPlaylist bool `yaml:"repeat_playlist"`
}

type musicSettings struct {
	RepeatSong     bool `yaml:"repeat_song"`
	RepeatPlaylist bool `yaml:"repeat_playlist"`
	Shuffle        bool `yaml:"shuffle"`
}

type userSetting struct {
	UseDB bool `yaml:"use_db"`
}

type TempSetting struct {
	StartDir      string   `yaml:"start_dir"`
	ShowIcons     bool     `yaml:"show_icons"`
	ShowHidden    bool     `yaml:"show_hidden"`
	Exclude       []string `yaml:"exclude"`
	PlayOnly      []string `yaml:"playonly"`
	Include       []string `yaml:"include"`
	Renderer      string   `yaml:"renderer"`
	EnableLogging bool     `yaml:"enable_logging"`
}

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
func Parse(temp *TempSetting) (Config, error) {
	var config Config
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

	if temp.StartDir != "" {
		config.General.StartDir = temp.StartDir
	}
	config.Temp = *temp

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
func (parser ConfigParser) readConfigFile(path string) (Config, error) {
	config := parser.getDefaultConfig()
	data, err := os.ReadFile(path)
	if err != nil {
		return config, configError{parser: parser, configDir: path, err: err}
	}

	err = yaml.Unmarshal((data), &config)
	return config, err
}

// getDefaultConfig returns the default config for the application.
func (parser ConfigParser) getDefaultConfig() Config {
	return Config{

		General: generalSettings{
			showIcons:     true,
			StartDir:      "~/Music/",
			EnableLogging: true,
		},
		Player: playerSettings{
			Shuffle:        true,
			RepeatPlaylist: true,
		},
		Music: musicSettings{
			RepeatSong:     false,
			RepeatPlaylist: true,
			Shuffle:        true,
		},
		Renderer: "raw",
		User: userSetting{
			UseDB: false,
		},
	}
}
