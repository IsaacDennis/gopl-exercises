package exercises

import "golang.org/x/net/html"

func Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	} else {
		GetLink(&links, n)
		if n.NextSibling != nil {
			links = Visit(links, n.NextSibling)
		}
		if n.FirstChild != nil {
			links = Visit(links, n.FirstChild)
		}
		return links
	}
}

func GetLink(links *[]string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				*links = append(*links, a.Val)
			}
		}
	}
}
