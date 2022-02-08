package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thechadgod/snoopy/helper"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace your IP",
	Long: `
Gets track IP address and location that is associated with it.
If you do not provide an IP address then it will use your 
public IP address.
`,
	Run: func(cmd *cobra.Command, args []string) {

				ipAddress := cmd.Flag("ip").Value.String()

		if ipAddress == "" {

			// get your public ip address

			publicIp, err := helper.GetPublicIP()

			if err != nil {
				panic(err)
			}

			response := fmt.Sprintf("IP : %s\nLocation: %s\nISP: %s", publicIp.Address, publicIp.Location, publicIp.ISP)

			fmt.Println(response)

		} else {

			// get the targets ip address

			targetIp, err := helper.TraceIP(ipAddress)

			if err != nil {
				panic(err)
			}

			targetResponse := fmt.Sprintf("IP : %s\nLocation: %s\nISP: %s", targetIp.Address, targetIp.Location, targetIp.ISP)

			fmt.Println(targetResponse)
		}

	},


}

func init() {
	rootCmd.AddCommand(traceCmd)
	traceCmd.Flags().StringP("ip", "i", "", "IP address to trace")

}
