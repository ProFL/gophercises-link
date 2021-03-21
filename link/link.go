package link

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLinksFromRoot(rootNode *html.Node) []Link {
	nodes := HTMLNodeMultimatchDFS(rootNode, func(node *html.Node) bool {
		return node.Type == html.ElementNode && node.Data == "a"
	})
	links := make([]Link, len(nodes))
	for i, node := range nodes {
		links[i] = ParseHTMLANode(node)
	}
	return links
}

func ParseHTMLANode(node *html.Node) (link Link) {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}

	childTextNodes := HTMLNodeMultimatchDFS(node, func(node *html.Node) bool {
		return node.Type == html.TextNode
	})
	texts := make([]string, len(childTextNodes))
	for _, textNode := range childTextNodes {
		texts = append(texts, strings.TrimSpace(textNode.Data))
	}
	link.Text = strings.TrimSpace(strings.Join(texts, " "))

	return link
}
