package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thechadgod/snoopy/helper"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace your IP",
	Long:  `Gets your IP address and location that is associated with it.`,
	Run: func(cmd *cobra.Command, args []string) {

		publicIp, err := helper.GetPublicIP()

		if err != nil {
			panic(err)
		}

		response := fmt.Sprintf("IP : %s\nLocation: %s\nISP: %s", publicIp.Address, publicIp.Location, publicIp.ISP)

		fmt.Println(response)

	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
