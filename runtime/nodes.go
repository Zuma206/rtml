package runtime

import (
	"strings"

	"golang.org/x/net/html"
)

type Handler func(runtime *Runtime, node *html.Node) error
type Handlers map[html.NodeType]Handler

var handlers Handlers

func init() {
	handlers = Handlers{
		html.DocumentNode: handleChildren,
		html.DoctypeNode:  handleDoctype,
		html.ElementNode:  handleElement,
		html.TextNode:     handleTextNode,
	}
}

func handleChildren(runtime *Runtime, node *html.Node) error {
	for child := range node.ChildNodes() {
		if err := runtime.eval(child); err != nil {
			return err
		}
	}
	return nil
}

func handleDoctype(runtime *Runtime, doctype *html.Node) error {
	if err := runtime.print("<!DOCTYPE ", doctype.Data, ">"); err != nil {
		return err
	}
	return nil
}

func handleElement(runtime *Runtime, element *html.Node) error {
	if isVoid, ok := voidElements[element.Data]; ok && isVoid {
		return handleVoidElement(runtime, element)
	}
	return handleRegularElement(runtime, element)
}

func handleVoidElement(runtime *Runtime, element *html.Node) error {
	return handleOpeningTag(runtime, element, true)
}

func handleRegularElement(runtime *Runtime, element *html.Node) error {
	if err := handleOpeningTag(runtime, element, false); err != nil {
		return err
	}
	if err := handleChildren(runtime, element); err != nil {
		return err
	}
	if err := handleClosingTag(runtime, element); err != nil {
		return err
	}
	return nil
}

func handleOpeningTag(runtime *Runtime, element *html.Node, selfClosing bool) error {
	if err := runtime.print("<", element.Data); err != nil {
		return err
	}
	for _, attr := range element.Attr {
		if err := runtime.print(" ", attr.Key, "=", "\"", attr.Val, "\""); err != nil {
			return err
		}
	}
	if selfClosing {
		return runtime.print("/>")
	}
	return runtime.print(">")
}

func handleClosingTag(runtime *Runtime, element *html.Node) error {
	return runtime.print("</", element.Data, ">")
}

func handleTextNode(runtime *Runtime, textNode *html.Node) error {
	return runtime.print(strings.TrimSpace(textNode.Data))
}
