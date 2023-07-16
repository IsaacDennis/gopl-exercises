package exercises

import (
	"golang.org/x/net/html"
	"unicode"
)

func GetTextContents(n *html.Node) []string {
	contents := []string{}
	if n.Type == html.TextNode {
		parentTag := n.Parent.Data
		if parentTag != "script" && parentTag != "style" {
			for _, char := range n.Data {
				if !unicode.IsSpace(char) {
					contents = append(contents, n.Data)
					break
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		contents = append(contents, GetTextContents(c)...)
	}
	return contents
}
