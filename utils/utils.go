package utils

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func AbsVal(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func GetDateDiffInDays(t int) int {
	return int(time.
		Since(time.
			Unix(int64(t), 0)).
		Hours() / 24)
}

func GenerateSeperatorString(w int) string {
	sep := ""
	for x := 0; x <= (98*w)/100; x++ {
		sep += "*"
	}
	return sep
}

func beautify(doc *html.Node) string {
	data := ""
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.TextNode {
			if node.Parent.Data == "code" {
				if len(strings.Split(node.Data, "\n")) > 1 {
					data += "\n"
					for _, s := range strings.Split(node.Data, "\n") {
						data += fmt.Sprintf("\t[yellow]%s\n", s)
					}
					data += "[-]"
				} else {
					data += fmt.Sprintf("[yellow]%s[-]", node.Data)
				}
			} else {
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
func BeautifyHtmlText(text string) string {
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		return "Something went wrong while parsing the ansewr"
	}
	return beautify(doc)
}
