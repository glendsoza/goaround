package utils

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// Helper function to get the difference between unix date and current date in number of days
// Used in answer template
func GetDateDiffInDays(t int) int {
	return int(time.
		Since(time.
			Unix(int64(t), 0)).
		Hours() / 24)
}

// Function to generate seperator string in anwser template
// Idea behind thsi function is to generate string with length equal to width of the terminal
// Coud not figure out how to do so this server as placeholder for now
func GenerateSeperatorString(w int) string {
	sep := ""
	for x := 0; x <= w; x++ {
		sep += "*"
	}
	return sep
}

// Beutify the html text
func beautify(doc *html.Node) string {
	data := ""
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.TextNode {
			if node.Parent.Data == "code" {
				// in case of multi line code block we want to add \t to the begining of each line
				// and highlight it with yello
				if len(strings.Split(node.Data, "\n")) > 1 {
					data += "[yellow]\n"
					for _, s := range strings.Split(node.Data, "\n") {
						data += fmt.Sprintf("\t%s\n", s)
					}
					data += "[-]"
				} else {
					// in case of single line code block add the color
					data += fmt.Sprintf("[yellow]%s[-]", node.Data)
				}
			} else {
				// append the rest of the text without any highlight color
				data += node.Data
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	return data
}

// Get the formatted text from the anwer body+
func BeautifyHtmlText(text string) string {
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		return "Something went wrong while parsing the ansewr"
	}
	return beautify(doc)
}
