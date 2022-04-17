/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var Tags string

var rootCmd = &cobra.Command{
	Use:   "goaround",
	Short: "A Tool to query stackoverflow via terminal",
	Long:  `A Tool to query stackoverflow via terminal`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
