package exercises

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	test := `
<html>
<body>
<h1>Heading 1</h1>
<h2>Heading 2</h2>
<div>
<h3>Heading 3</h3>
<img src="test.img"></img>
</div>
<script>Test</script>
</body>
</html>
`
	reader := strings.NewReader(test)
	doc, err := html.Parse(reader)
	if err != nil {
		t.Errorf("Error while parsing HTML: %s", err)
	}
	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	t.Log(images)
	t.Log(headings)
}
