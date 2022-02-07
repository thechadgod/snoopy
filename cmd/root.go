package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cli-go",
		Short: "I am a cli-go",
		Long:  `Kinda like a cli-go, but not really.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
