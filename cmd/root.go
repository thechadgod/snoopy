package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "snoopy",
		Short: "snoop around people on the internet",
		Long:  `
Track other people on the internet and gather data on them.
Automate opsec and osint with snoopy.
`,
	}
)

func Execute()  {
		err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
