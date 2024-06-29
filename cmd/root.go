/*
Copyright © 2024 achyut koirala <axyut@outlook.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/axyut/playgo/internal/app"
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
	Version:   "0.1.3",
	ValidArgs: []string{"."},
	Args:      cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		config, err := config.Parse()
		if err != nil {
			log.Fatal(err)
		}
		app.StartPlaygo(config)
	},
	Example: `playgo [directory]`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)

	rootCmd.PersistentFlags().String("start-dir", ".", "Starting directory for Playgo")
	rootCmd.PersistentFlags().String("include", "", "File/s to play")
	rootCmd.PersistentFlags().String("exclude", "", "File/s to ignore")
	rootCmd.PersistentFlags().String("renderer", "raw", "Application Renderer")
	rootCmd.PersistentFlags().Bool("show-icons", true, "Show icons")
}
