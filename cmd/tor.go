package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thechadgod/snoopy/helper"
)

var torCmd = &cobra.Command{
	Use:   "tor",
	Short: "Hide your IP",
	Long: `
Anonymize your IP address by using Tor.
Tor will use onion routing to route your IP address through the Tor network.
`,
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Tor",
	Long: `
Start Tor.
`,
	Run: func(cmd *cobra.Command, args []string) {

		// start tor
		helper.StartTor()

	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Tor",
	Long: `
Stop Tor.
`,
	Run: func(cmd *cobra.Command, args []string) {

		// stop tor
		helper.StopTor()

	},
}

var statusCmd = &cobra.Command{

	Use:   "status",
	Short: "Check Tor status",
	Long: `
Check Tor status.
`,
	Run: func(cmd *cobra.Command, args []string) {

		// check tor status
		helper.StatusTor()

	},
}

func init() {
	torCmd.AddCommand(startCmd)
	torCmd.AddCommand(stopCmd)
	torCmd.AddCommand(statusCmd)

	rootCmd.AddCommand(torCmd)
}
