package exercises

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestGetElementOcurrences(t *testing.T) {
	htmlTest := `
<html>
<head>
</head>
<body>
<div></div>
<div></div>
<p></p>
</body>
</html>
`
	reader := strings.NewReader(htmlTest)
	doc, err := html.Parse(reader)
	ocur := make(map[string]int)
	if err != nil {
		t.Errorf("Error while parsing HTML: %s", err)
	}
	GetElementOcurrences(doc, ocur)
	t.Log(ocur)
}
