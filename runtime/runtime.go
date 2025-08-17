package runtime

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Runtime struct {
	Output io.Writer
}

func New() *Runtime {
	return &Runtime{}
}

func (runtime *Runtime) RunCode(code io.Reader) error {
	document, err := html.Parse(code)
	if err != nil {
		return err
	}
	return runtime.eval(document)
}

type Handler func(runtime *Runtime, node *html.Node) error
type Handlers map[html.NodeType]Handler

var handlers Handlers = Handlers{}

func (runtime *Runtime) eval(node *html.Node) error {
	handler, ok := handlers[node.Type]
	if !ok {
		return fmt.Errorf("cannot handle node of type %d", node.Type)
	}
	return handler(runtime, node)
}
