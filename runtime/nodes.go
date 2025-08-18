package runtime

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

type Handler func(runtime *Runtime, node *html.Node) error

var handlers map[html.NodeType]Handler
var special map[string]Handler

func init() {
	handlers = map[html.NodeType]Handler{
		html.DocumentNode: handleChildren,
		html.DoctypeNode:  handleDoctype,
		html.ElementNode:  handleElement,
		html.TextNode:     handleTextNode,
	}

	special = map[string]Handler{
		"script":   handleScript,
		"template": handleSkip,
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
	if specialHandler, isSpecial := special[element.Data]; isSpecial {
		return specialHandler(runtime, element)
	}
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

func handleScript(runtime *Runtime, scriptElement *html.Node) error {
	var buffer bytes.Buffer
	output := runtime.Stream
	runtime.Stream = &buffer
	if err := handleChildren(runtime, scriptElement); err != nil {
		return err
	}
	runtime.Stream = output
	if _, err := runtime.VM.RunString(buffer.String()); err != nil {
		return err
	}
	return nil
}

func handleSkip(_ *Runtime, _ *html.Node) error {
	return nil
}
