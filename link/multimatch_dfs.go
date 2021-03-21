package link

import "golang.org/x/net/html"

func HTMLNodeMultimatchDFS(node *html.Node, targetNodeFn func(node *html.Node) bool) []*html.Node {
	return _HTMLNodeMultimatchDFS(node, node, targetNodeFn)
}

func _HTMLNodeMultimatchDFS(rootNode *html.Node, node *html.Node, targetNodeFn func(node *html.Node) bool) []*html.Node {
	var nodes []*html.Node
	if targetNodeFn(node) {
		nodes = append(nodes, node)
	}
	if node.FirstChild != nil {
		nodes = append(nodes, _HTMLNodeMultimatchDFS(rootNode, node.FirstChild, targetNodeFn)...)
	}
	if rootNode != node && node.NextSibling != nil {
		nodes = append(nodes, _HTMLNodeMultimatchDFS(rootNode, node.NextSibling, targetNodeFn)...)
	}
	return nodes
}
