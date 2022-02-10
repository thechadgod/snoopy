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

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change address or identity.",
	Long: `
Change your mac address or tor identity.
	`,
}

var changeIdCmd = &cobra.Command{
	Use:   "id",
	Short: "Change Tor identity",
	Long: `
Changes your tor identity.	
`,
	Run: func(cmd *cobra.Command, args []string) {

		// change tor id
		helper.ChangeID()

	},
}

var changeMacCmd = &cobra.Command{
	Use:   "mac",
	Short: "Change mac address",
	Long: `
Change your mac address.	
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// change mac address
		helper.ChangeMac()

	},
}

func init() {

	changeCmd.AddCommand(changeIdCmd)
	changeCmd.AddCommand(changeMacCmd)

	torCmd.AddCommand(startCmd)
	torCmd.AddCommand(stopCmd)
	torCmd.AddCommand(statusCmd)
	torCmd.AddCommand(changeCmd)

	rootCmd.AddCommand(torCmd)
}
