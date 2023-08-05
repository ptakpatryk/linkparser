package parser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseHtmlLinks(r io.Reader) []Link {
	fmt.Println("--- Parsing HTML")
	doc, _ := html.Parse(r)

	var links []Link

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			text := n.FirstChild.Data
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, Link{Href: a.Val, Text: text})
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links
}

func isLink(tag []byte) bool {
	return len(tag) == 1 && tag[0] == 'a'
}
