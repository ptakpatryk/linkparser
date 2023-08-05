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
		if isLink(n) {
			links = append(links, parseLink(n))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links
}

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func parseLink(ln *html.Node) (linkItem Link) {
	text := ln.FirstChild.Data
	for _, a := range ln.Attr {
		if a.Key == "href" {
			linkItem = Link{Href: a.Val, Text: text}
		}
	}

	return linkItem
}
