/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"hep/internal/inspire"
	"github.com/spf13/cobra"
)

// literatureCmd represents the literature command
var literatureCmd = &cobra.Command{
	Use:   "literature",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		q, err := cmd.Flags().GetString("query")
		if err != nil {
			panic(err)
		}
		omitAbstract, err := cmd.Flags().GetBool("no-abstract")
		if err != nil {
			panic(err)
		}

		inspire.QueryResults(q, omitAbstract)
	},
}

func init() {
	literatureCmd.Flags().StringP("query", "q", "", "Most recient papers of following authors")
	literatureCmd.Flags().Bool("no-abstract", false, "Most recient papers of following authors")
	rootCmd.AddCommand(literatureCmd)
}
