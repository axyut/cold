/*
Copyright © 2024 achyut koirala <axyut.github.io>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/axyut/playgo/internal/app"
	"github.com/axyut/playgo/internal/booTea"
	"github.com/axyut/playgo/internal/config"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "playgo",
	Short: "A CLI Music Player",
	Long: `A CLI Music Player that plays mp3 files from a directory, defaults to the current directory,
if not found any music files, plays from ~/Music/. It provides a simple interface to play, pause,
skip, and repeat songs.`,
	Version: "0.1.3",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// If logging is enabled, logs will be output to debug.log.
		// if enableLogging {
		// 	f, err := tea.LogToFile("debug.log", "debug")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	defer func() {
		// 		if err = f.Close(); err != nil {
		// 			log.Fatal(err)
		// 		}
		// 	}()
		// }

		setting := getTempSettings(cmd, args)
		config, err := config.Parse(setting)
		if err != nil {
			log.Fatal(err)
		}
		if config.Renderer == "tea" {
			booTea.RunBubbleTUI()
		} else {
			app.StartPlaygo(config)
		}
	},
	Example: `playgo # no commands defaults to config's start directory
playgo . # if no audio files, defaults to ~/Music/
playgo ~/Music -e a.mp3 -e b.mp3`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(updateCmd)

	rootCmd.PersistentFlags().StringArrayP("exclude", "e", []string{}, "File/s to ignore while playing files in directory")
	rootCmd.PersistentFlags().StringArrayP("include", "i", []string{}, "Include File/s to play with files in directory")
	rootCmd.PersistentFlags().StringArrayP("playonly", "p", []string{}, "Only File/s to play")
	rootCmd.PersistentFlags().StringP("renderer", "r", "raw", "Application Renderer [raw, tea]")
	rootCmd.PersistentFlags().Bool("icons", true, "Show icons [true/false]")
	rootCmd.PersistentFlags().Bool("hidden", false, "Play Hidden Files [true/false]")
	rootCmd.PersistentFlags().Bool("logging", false, "Enable logging player [true/false]")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getTempSettings(cmd *cobra.Command, args []string) *config.TempSetting {
	startDir := ""
	if len(args) > 0 {
		// fmt.Println("Start Dir: ", args[0])
		startDir = args[0]

	}

	enableLogging, err := cmd.Flags().GetBool("logging")
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := cmd.Flags().GetString("renderer")
	// log.Println("Renderer: ", renderer)
	if err != nil {
		log.Fatal(err)
	}

	showIcons, err := cmd.Flags().GetBool("icons")
	if err != nil {
		log.Fatal(err)
	}
	showHidden, err := cmd.Flags().GetBool("hidden")
	if err != nil {
		log.Fatal(err)
	}

	exclude, err := cmd.Flags().GetStringArray("exclude")
	if err != nil {
		log.Fatal(err)
	}

	include, err := cmd.Flags().GetStringArray("include")
	if err != nil {
		log.Fatal(err)
	}

	playOnly, err := cmd.Flags().GetStringArray("playonly")
	if err != nil {
		log.Fatal(err)
	}
	return &config.TempSetting{
		StartDir:      startDir,
		EnableLogging: enableLogging,
		Renderer:      renderer,
		ShowIcons:     showIcons,
		ShowHidden:    showHidden,
		Exclude:       exclude,
		Include:       include,
		PlayOnly:      playOnly,
	}
}
