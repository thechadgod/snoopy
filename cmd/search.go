package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thechadgod/snoopy/internal/helper/search"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Use the web for osint searches",
	Long: `
Use the web for osint searches.
`,
}

var usernameCmd = &cobra.Command{
	Use:     "username",
	Aliases: []string{"u"},
	Short:   "Search for usernames",
	Args:    cobra.MinimumNArgs(1),
	Long: `
Search for usernames.
`, Run: func(cmd *cobra.Command, args []string) {

		username := args[0]

		if priority, err := cmd.Flags().GetBool("priority"); err == nil {
			search.Username(username, priority)
		} else {
			fmt.Println("Error: ", err)

		}

	},
}

func init() {

	usernameCmd.Flags().BoolP("priority", "p", false, "Search only the priority sites")

	searchCmd.AddCommand(usernameCmd)

	rootCmd.AddCommand(searchCmd)

}
