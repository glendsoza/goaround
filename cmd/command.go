/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"goaround/executor"
	"goaround/ui"
	"log"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Command to execute",
	Long:  `Command to execute`,
	Run: func(cmd *cobra.Command, args []string) {
		errorString, executable := executor.Execute(args[0])
		// if error string is not empty then the command failed
		if errorString != "" {
			var input string
			fmt.Println(errorString)
			fmt.Println("Do you want to display Stackoverflow results? (y/n)")
			fmt.Scanln(&input)
			if input == "y" || input == "Y" {
				if err := ui.Run(errorString, executable); err != nil {
					log.Fatal(err)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(commandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
