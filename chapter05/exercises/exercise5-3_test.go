package exercises

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestGetTextContents(t *testing.T) {
	test := `
<html>
<body>
<p>Text 1</p>
<p>Text 2</p>
<div>
<p>Text 3</p>
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
	result := GetTextContents(doc)
	t.Log(result)
}
