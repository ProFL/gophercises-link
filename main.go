package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var sampleHtml = `
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>`

func HTMLNodeMultimatchDFS(node *html.Node, targetNodeFn func(node *html.Node) bool, keepDigging bool) []*html.Node {
	var nodes []*html.Node
	if targetNodeFn(node) {
		nodes = append(nodes, node)
		if !keepDigging {
			return nodes
		}
	}
	if node.FirstChild != nil {
		nodes = append(nodes, HTMLNodeMultimatchDFS(node.FirstChild, targetNodeFn, keepDigging)...)
	}
	if node.NextSibling != nil {
		nodes = append(nodes, HTMLNodeMultimatchDFS(node.NextSibling, targetNodeFn, keepDigging)...)
	}
	return nodes
}

func isLinkNode(node *html.Node) bool {
	return node.Type == html.ElementNode && node.Data == "a"
}

func main() {
	reader := strings.NewReader(sampleHtml)
	rootNode, err := html.Parse(reader)
	if err != nil {
		log.Fatalln("Failed to parse HTML", err.Error())
	}

	nodes := HTMLNodeMultimatchDFS(rootNode, isLinkNode, false)
	links := make([]Link, len(nodes))
	for i, node := range nodes {
		links[i] = Link{Text: ""}
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links[i].Href = attr.Val
				break
			}
		}
		childTextNodes := HTMLNodeMultimatchDFS(node, func(node *html.Node) bool {
			return node.Type == html.TextNode
		}, true)
		texts := make([]string, len(childTextNodes))
		for _, textNode := range childTextNodes {
			texts = append(texts, strings.TrimSpace(textNode.Data))
		}
		links[i].Text = strings.TrimSpace(strings.Join(texts, " "))
	}
	fmt.Printf("%+v\n", links)
}
