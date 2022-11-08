package webpage_requisites_go

import (
	"sort"
	"strings"
	"testing"
)

func TestHtmlRequisites(t *testing.T) {
	html := `<!DOCTYPE html>
	<html>
	<head>
		<style>* { background: url('a.jpg'); }</style>
		<link rel="stylesheet" href="b.css">
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
	sort.SliceStable(urls, func(i, j int) bool {
		return urls[i].String() < urls[j].String()
	})

	if len(urls) != 4 {
		t.Errorf("not the right number of urls found")
	}

	if urls[0].String() != "a.jpg" {
		t.Errorf("url a.jpg not found")
	}
	if urls[1].String() != "b.css" {
		t.Errorf("url b.css not found")
	}
	if urls[2].String() != "c.js" {
		t.Errorf("url c.js not found")
	}
	if urls[3].String() != "d.png" {
		t.Errorf("url d.png not found")
	}
}
