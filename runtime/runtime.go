package runtime

import (
	"fmt"
	"io"

	"github.com/dop251/goja"
	"golang.org/x/net/html"
)

type Runtime struct {
	Output io.Writer
	vm     *goja.Runtime
}

func New() *Runtime {
	return &Runtime{
		vm: goja.New(),
	}
}

func (runtime *Runtime) RunCode(code io.Reader) error {
	document, err := html.Parse(code)
	if err != nil {
		return err
	}
	return runtime.eval(document)
}

func (runtime *Runtime) eval(node *html.Node) error {
	handler, ok := handlers[node.Type]
	if !ok {
		return fmt.Errorf("cannot handle node of type %d", node.Type)
	}
	return handler(runtime, node)
}

func (runtime *Runtime) print(args ...any) error {
	_, err := fmt.Fprint(runtime.Output, args...)
	return err
}
