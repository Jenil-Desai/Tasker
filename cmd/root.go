package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Tasker",
	Short: "Tasker Command Center - Mission Control Interface",
	Long:  `Welcome to Tasker Command Center - Your terminal-based mission control interface. Deploy, track, and accomplish tasks with military precision. Awaiting your command, operator.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
