package webpage_requisites_go

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/url"
	"strings"
)

// GetHtmlRequisites might return duplicates
func GetHtmlRequisites(html io.Reader) ([]*url.URL, error) {
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return nil, err
	}

	urls := make([]*url.URL, 0)

	doc.Find("link").Each(func(_ int, selection *goquery.Selection) {
		rel, exists := selection.Attr("rel")
		if !exists {
			return
		}
		if strings.ToLower(rel) != "stylesheet" {
			return
		}
		href, exists := selection.Attr("href")
		if !exists {
			return
		}
		u, err := url.Parse(href)
		if err != nil {
			return
		}
		urls = append(urls, u)
	})

	doc.Find("img").Each(func(_ int, selection *goquery.Selection) {
		src, exists := selection.Attr("src")
		if !exists {
			return
		}
		if strings.HasPrefix(strings.ToLower(src), "data:") {
			return // skip embedded data
		}
		u, err := url.Parse(src)
		if err != nil {
			return
		}
		urls = append(urls, u)
	})

	doc.Find("script").Each(func(_ int, selection *goquery.Selection) {
		src, exists := selection.Attr("src")
		if !exists {
			return
		}
		u, err := url.Parse(src)
		if err != nil {
			return
		}
		urls = append(urls, u)
	})

	doc.Find("style").Each(func(_ int, selection *goquery.Selection) {
		cssUrls, err := GetCssRequisites(selection.Text())
		if err != nil {
			return
		}
		urls = append(urls, cssUrls...)
	})

	return urls, nil
}
