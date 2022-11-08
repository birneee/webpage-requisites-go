package webpage_requisites_go

import (
	"github.com/gorilla/css/scanner"
	"net/url"
	"regexp"
	"strings"
)

var CSS_URL_REGEXP *regexp.Regexp = regexp.MustCompile(`^url\(['"]?(.*?)['"]?\)$`)

// GetCssRequisites might return duplicates
// inspired by https://github.com/cornelk/goscrape
func GetCssRequisites(css string) ([]*url.URL, error) {
	urls := make([]*url.URL, 0)
	s := scanner.New(css)
	for {
		token := s.Next()
		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}
		if token.Type != scanner.TokenURI {
			continue
		}
		match := CSS_URL_REGEXP.FindStringSubmatch(token.Value)
		if match == nil {
			continue
		}
		src := match[1]
		if strings.HasPrefix(strings.ToLower(src), "data:") {
			continue // skip embedded data
		}
		u, err := url.Parse(src)
		if err != nil {
			continue
		}
		urls = append(urls, u)
	}

	return urls, nil
}
