# Go Around : A tool to Query stackoverflow via terminal

![Demo](goaround.gif)

## Overview

Go around uses the stackoverflow API to get the answers for the given query and display them in terminal.

## Installation

Download the latest binary corresponding to your platform from [releases](https://github.com/glendsoza/goaround/releases/tag/v0.5) page

## Usage

### Querying the API directly

```bash
./goaround -q "<your query>"
```

```bash
./goaround -q "Python upper case a string"
```

```bash
export STACKOVERFLOW_APP_KEY="<your app key>"
export STACKOVERFLOW_PAGE_SIZE=50
./goaround -q "<your query>"
```

```bash
export STACKOVERFLOW_APP_KEY="<your app key>"
export STACKOVERFLOW_PAGE_SIZE=50
./goaround -q "Python upper case a string"
```

To get more accurate results you can pass the tags as comma separated values, results containing at least one of the tags will be shown

```bash
./goaround -q "<your query>" -t "<comma seperated values>"
```

```bash
./goaround -q "Python upper case a string" -t "python,python3"
```

### Using `goaround` as wrapper to run other programs (currently only supports go and python)

`goaround` can be used to capture the `Stderr` of other porgrams and query the stackoverflow API with the erorr generated and display the results.

```bash
./goaround -p "<your command>"
```

```bash
./goaround -p "go run main.go"
```

```bash
./goaround -p "python main.py"
```

Few things to note while using `goaround` as wrapper:

- Stdout of the command will be displayed in the terminal in real time
- Only the error pushed to `Stderr` will be used to query the stackoverflow API
- In this mode tags provided via `-t` will be ignored and name of the executable in the command will be taken as tag

### Navigating through the results

- In questions screen use the Mouse Scroll or Arrow Keys to go up and down
- Use the Enter key to open the answer
- Use the Backspace key to go back from answers to the questions screen
- Use the `Ctrl + c` key to exit the application
- Hot keys with sepcial function in each screen will be displayed on the bottom of the screen

## Configuration

Following environment can be used to configure the tool.

| Name                    | Required | Default | Description                                                                                                                                                                                                                                                                                      |
| ----------------------- | -------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| STACKOVERFLOW_APP_KEY   | No       | -       | There is a limit on the number of calls that can be made to the stackoverflow api, by default 300 requests can be made per day, by providing API key this can be increased to 10000 requests per day. App can be registered [here](https://stackapps.com/apps/oauth/register) to get the App key |
| STACKOVERFLOW_PAGE_SIZE | No       | 25      | Number of questions displayed in the terminal. By default its 25 and can be set upto 100                                                                                                                                                                                                         |

## Contributing

Following Features are planned to be added in the future but any help is welcome!

- Make results more accurate and relevant
- Provide button to copy the code from the answers to the clipboard
