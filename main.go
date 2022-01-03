package main

import (
	"flag"
	"fmt"
	"goaround/executor"
	"goaround/ui"
	"log"
)

var (
	command,
	query,
	tags string
)

func main() {

	flag.StringVar(&query, "q", "", "Query to search")
	flag.StringVar(&tags, "t", "", "List of command seperated tags to narrow down the search")
	flag.StringVar(&command, "p", "", "Command to execute")
	flag.Parse()
	if query == "" && command == "" {
		log.Fatal("Please pass either a query or a command to execute")
	}
	// check if any command is provided
	if command != "" {
		errorString, executable := executor.Execute(command)
		// if error string is not empty then the command failed
		if errorString != "" {
			var input string
			fmt.Println(errorString)
			fmt.Println("Do you want to display Stackoverflow results? (y/n)")
			fmt.Scanln(&input)
			if input == "y" || input == "Y" {
				query = errorString
				tags = executable
			} else {
				return
			}
		} else {
			// if error string is empty then the command ran successfully
			return
		}
	}
	// Initialize the ui depending on if q is provided or command provided via p failed
	if err := ui.Run(query, tags); err != nil {
		log.Fatal(err)
	}
}
