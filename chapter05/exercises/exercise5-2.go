package exercises

import "golang.org/x/net/html"

func GetElementOcurrences(n *html.Node, ocur map[string]int) {
	if n.Type == html.ElementNode {
		ocur[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		GetElementOcurrences(c, ocur)
	}
}
