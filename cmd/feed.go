/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"hep/internal/inspire"
	"github.com/spf13/cobra"
)

// feedCmd represents the feed command
var feedCmd = &cobra.Command{
	Use:   "feed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		OnlyFollowers, err := cmd.Flags().GetBool("following")
		OmitAbstract, err := cmd.Flags().GetBool("no-abstract")
		if err != nil {
			panic(err)
		}
		inspire.FeedMenu(OnlyFollowers, OmitAbstract)
	},
}

/*
 * TODO Add flag --limit
 * TODO Add flag --no-abstract 
 * TODO Add flag --pdf-export 
 */
func init() {
	feedCmd.Flags().BoolP("following", "f", false, "Most recient papers of following authors")
	feedCmd.Flags().Bool("no-abstract", false, "Omits abstracts")
	rootCmd.AddCommand(feedCmd)
}
