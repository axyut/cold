package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Playgo",
	Long:  `Update Playgo to the latest version.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateCommand := exec.Command(
			"bash",
			"-c",
			"curl -sfL https://raw.githubusercontent.com/axyut/playgo/master/install.sh | sh",
		)

		// check if go is installed, if yes, use go install @latest

		// check if there are releases, if yes, use the latest release
		// find os, arch and use the correct binary

		// check if there is install script, if yes, use the install script
		updateCommand.Stdin = os.Stdin
		updateCommand.Stdout = os.Stdout
		updateCommand.Stderr = os.Stderr

		err := updateCommand.Run()
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	},
}
