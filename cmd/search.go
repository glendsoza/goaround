/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"goaround/ui"
	"log"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Query to search",
	Long:  `Query to search`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ui.Run(args[0], Tags); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&Tags, "tags", "t", "", "Comma Seprated tags i.e go,pyhon")
}
