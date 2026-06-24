/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"hep/internal/inspire"
)

// citeCmd represents the cite command
var citeCmd = &cobra.Command{
	Use:   "cite",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		id, err := strconv.Atoi(args[0])

		if err != nil {
			panic(err)
		}

		citation := inspire.FetchCitation(id)
		fmt.Println(citation)
	},
}

func init() {
	rootCmd.AddCommand(citeCmd)
}
