package htmlquery

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type HTML struct {
}

func ReadFromString(s string) *html.Node {
	doc, err := htmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err.Error())
	}
	return doc
}

func ReadFromByte(b []byte) *html.Node{
	return ReadFromString(string(b))
}
