package runtime

import "golang.org/x/net/html"

type Template struct {
	*html.Node
}

type Templates map[string]Template

func (templates Templates) Scan(node *html.Node) {
	if node.Type != html.ElementNode && node.Type != html.DocumentNode {
		return
	}
	if node.Type == html.DocumentNode || node.Data != "template" {
		for child := range node.ChildNodes() {
			templates.Scan(child)
		}
		return
	}
	for _, attr := range node.Attr {
		if attr.Key == "id" {
			templates[attr.Val] = Template{node}
			return
		}
	}
}
