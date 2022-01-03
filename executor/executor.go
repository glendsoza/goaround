package executor

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var PYTHON_EXEPECTED_ERRORS_REGEX = regexp.MustCompile(`KeyboardInterrupt|SystemExit|GeneratorExit`)

func Execute(command string) (string, string) {
	// Split the command into tokens
	commandTokens := strings.Split(command, " ")
	// grab the executable name
	executable := commandTokens[0]
	cmd := exec.Command(executable, commandTokens[1:]...)
	var errorBuff bytes.Buffer
	// print the stdout
	cmd.Stdout = os.Stdout
	// Storing the error in buffer inf any
	cmd.Stderr = &errorBuff
	err := cmd.Start()
	// Exit if commands cannot be run
	if err != nil {
		log.Println("Please fix the following error")
		log.Fatal(err)
	}
	cmd.Wait()
	errorString := errorBuff.String()
	// If erorr string is not null then the command exited with the error
	if errorString != "" {
		fmt.Printf("\nError >>>>>>>>>>\n\n %s \n <<<<<<<<<< Error\n", errorString)
		// Depending on the executable return the error string
		switch executable {
		case "go":
			data := strings.Split(errorString, "\n")[0]
			if len(strings.Split(data, ": ")) > 1 {
				return strings.Join(strings.Split(data, ": ")[1:], " "), executable
			}
		case "python", "python3":
			// expected errors
			if !PYTHON_EXEPECTED_ERRORS_REGEX.MatchString(errorString) {
				data := strings.Split(errorString, "\n")
				return data[len(data)-2], executable
			}
		case "default":
			log.Println("Unrecogized executable")
		}
	}
	return "", ""
}
