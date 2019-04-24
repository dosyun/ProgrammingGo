package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := getHTMLDoc("https://golang.org/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "getHTMLDoc() error [%v]\n", err)
		return
	}

	nodes := ElementsByTagName(doc, "a", "div", "span")
	for _, node := range nodes {
		fmt.Println(node)
	}
}

func getHTMLDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return html.Parse(resp.Body)
}

func ElementsByTagName(n *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node

	var forEachNode func(n *html.Node, name ...string)
	forEachNode = func(n *html.Node, name ...string) {
		if n.Type == html.ElementNode && contains(n.Data, name...) {
			nodes = append(nodes, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, name...)
		}
	}
	forEachNode(n, name...)

	return nodes
}

func contains(name string, names ...string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}
	return false
}
