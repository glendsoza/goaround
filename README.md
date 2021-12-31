# Go Around : A tool to Query stackoverflow via terminal

![Demo](goaround.gif)

## Overview

Go around uses the stackoverflow API to get the answers for the given query and display them in terminal.

## Installation

Download the binary corresponding to your platform from releases page

## Usage

```bash
./goaround -q "<your query>"
```

With environment variable

```bash
export STACKOVERFLOW_APP_KEY="<your app key>"
export STACKOVERFLOW_PAGE_SIZE=50
./goaround -q "<your query>"
```

## Configuration

Following environment can be used to configure the tool.

| Name                    | Required | Default | Description                                                                                                                                                                                                                                                                                    |
| ----------------------- | -------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| STACKOVERFLOW_APP_KEY   | No       | -       | There is a limit on the number of calls that can be made to the stackoverflow api by default 300 requests can be made per day, by provinding API key this can be increased to 10000 requests per day. Apps can be registered here to get the App key https://stackapps.com/apps/oauth/register |
| STACKOVERFLOW_PAGE_SIZE | No       | 25      | Number of questions disaplayed in the terminal by default its 25 and can be set upto 100                                                                                                                                                                                                       |

## Contributing

Following Features are planned to be added in the future but any help is welcome!

- Make the tool similar to [Rebound](https://github.com/shobrook/rebound)
- Make the Answer in the terminal selectable (for seleting and copying the answer)
- Provide button to copy the code from the answers to the clipboard
