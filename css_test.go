package webpage_requisites_go

import (
	"sort"
	"testing"
)

func TestCssRequisites(t *testing.T) {
	css := `* { 
		background: url('a.jpg');
		cursor: url(b.cur);
		content: url("c.png");
		border-image: url('data:image/png;base64,...'); /* should be ignored */
		/* background-image: url('d.jpg'); should be ignored */
	}`
	urls, err := GetCssRequisites(css)
	if err != nil {
		t.Errorf("Failed to get requisites: %v", err)
	}
	sort.SliceStable(urls, func(i, j int) bool {
		return urls[i].String() < urls[j].String()
	})

	if len(urls) != 3 {
		t.Errorf("not the right number of urls found")
	}

	if urls[0].String() != "a.jpg" {
		t.Errorf("url a.jpg not found")
	}
	if urls[1].String() != "b.cur" {
		t.Errorf("url b.cur not found")
	}
	if urls[2].String() != "c.png" {
		t.Errorf("url c.png not found")
	}
}
