package exercises

import "golang.org/x/net/html"

func VisitLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link":
			links = append(links, GetAttr(n, "href"))
		case "img", "script":
			links = append(links, GetAttr(n, "src"))
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = VisitLinks(links, c)
	}
	return links
}

func GetAttr(n *html.Node, attr string) string {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val
		}
	}
	return ""
}
