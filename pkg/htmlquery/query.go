package htmlquery

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func QueryAll(doc *html.Node, expr string) []*html.Node {
	list, err := htmlquery.QueryAll(doc, expr)
	if err != nil {
		panic(err.Error())
	}
	return list
}

func FindOne(top *html.Node, expr string) *html.Node {
	node := htmlquery.FindOne(top, expr)
	return node
}

func FindOneText(top *html.Node,expr string) string{
	n := FindOne(top,expr)
	return InnerText(n)
}

func InnerText(n *html.Node) string {
	return htmlquery.InnerText(n)
}

func SelectAttr(n *html.Node, name string) (val string) {
	return htmlquery.SelectAttr(n, name)
}
