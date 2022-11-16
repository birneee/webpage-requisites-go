package webpage_requisites_go

import (
	"golang.org/x/exp/slices"
	"strings"
	"testing"
)

func TestHtmlRequisites(t *testing.T) {
	html := `<!DOCTYPE html>
	<html>
	<head>
		<style>* { background: url('a.jpg'); }</style>
		<link rel="stylesheet" href="b.css">
		<link rel="manifest" href="manifest.json">
		<link rel="shortcut icon" href="favicon.ico">
		<link href="f.css"> <!-- should be ignored -->
		<script src="c.js"></script>
	</head>
	<body>
		<img src="d.png">
		<img src="data:image/gif;base64,..."> <!-- should be ignored -->
		<a href="e.html">link</a> <!-- this is no requisite -->
	</body>
	</html>`
	urls, err := GetHtmlRequisites(strings.NewReader(html))
	if err != nil {
		t.Errorf("Failed to get requisites: %v", err)
	}

	var foundUrlStrings []string
	for _, u := range urls {
		foundUrlStrings = append(foundUrlStrings, u.String())
	}

	expectedUrls := []string{"a.jpg", "b.css", "c.js", "d.png", "manifest.json", "favicon.ico"}

	if len(urls) != len(expectedUrls) {
		t.Errorf("not the right number of urls found: found %d, exprected %d", len(foundUrlStrings), len(expectedUrls))
	}

	for _, expectedUrl := range expectedUrls {
		if !slices.Contains(foundUrlStrings, expectedUrl) {
			t.Errorf("url %s not found", expectedUrl)
		}
	}
}
