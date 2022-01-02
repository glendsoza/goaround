package main

import (
	"flag"
	"fmt"
	"goaround/api"
	"goaround/executor"
	"goaround/ui"
	"log"
)

func main() {
	flag.StringVar(&api.Query, "q", "", "Query to search")
	flag.StringVar(&api.Tags, "t", "", "List of command seperated tags to narrow down the search")
	flag.StringVar(&executor.Command, "p", "", "Command to execute")
	flag.Parse()
	if api.Query == "" && executor.Command == "" {
		log.Fatal("Please pass either a query or a command to execute")
	}
	// check if any command is provided
	if executor.Command != "" {
		errorString, executable := executor.Execute()
		// if error string is not empty then the command failed
		if errorString != "" {
			var input string
			fmt.Println(errorString)
			fmt.Println("Do you want to display Stackoverflow results? (y/n)")
			fmt.Scanln(&input)
			if input == "y" || input == "Y" {
				api.Query = errorString
				api.Tags = executable
			} else {
				return
			}
		} else {
			// if error string is empty then the command ran successfully
			return
		}
	}
	// Initialize the ui depending on if q is provided or command provided via p failed
	ui.InIt()
}
