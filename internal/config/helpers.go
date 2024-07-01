package config

import (
	"errors"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

// parsingError represents an error that occurred while parsing the config file.
type parsingError struct {
	err error
}

// Error represents an error that occurred while parsing the config file.
func (e parsingError) Error() string {
	return fmt.Sprintf("failed parsing config.yml: %v", e.err)
}

// createConfigFileIfMissing creates the config file if it doesn't exist.
func (parser ConfigParser) createConfigFileIfMissing(configFilePath string) error {
	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		newConfigFile, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			return err
		}

		defer newConfigFile.Close()
		return parser.writeDefaultConfigContents(newConfigFile)
	}

	return nil
}

// getDefaultConfigYamlContents returns the default config file contents.
func (parser ConfigParser) getDefaultConfigYamlContents() string {
	defaultConfig := parser.getDefaultConfig()
	yaml, _ := yaml.Marshal(defaultConfig)

	return string(yaml)
}

// Error returns the error message for when a config file is not found.
func (e configError) Error() string {
	return fmt.Sprintf(
		`Couldn't find a config.yml configuration file.
Create one under: %s
Example of a config.yml file:
%s
For more info, go to https://github.com/axyut/playgo
press q to exit.
Original error: %v`,
		path.Join(e.configDir, AppDir, ConfigFileName),
		e.parser.getDefaultConfigYamlContents(),
		e.err,
	)
}

// writeDefaultConfigContents writes the default config file contents to the given file.
func (parser ConfigParser) writeDefaultConfigContents(newConfigFile *os.File) error {
	_, err := newConfigFile.WriteString(parser.getDefaultConfigYamlContents())

	if err != nil {
		return err
	}

	return nil
}
