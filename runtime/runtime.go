package runtime

import (
	"fmt"
	"io"
	"os"

	"github.com/dop251/goja"
	"golang.org/x/net/html"
)

type Runtime struct {
	Stream io.Writer
	Log    io.Writer
	VM     *goja.Runtime
}

func New() *Runtime {
	runtime := &Runtime{
		Stream: os.Stdout,
		Log:    os.Stdout,
		VM:     goja.New(),
	}
	return runtime
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
	_, err := fmt.Fprint(runtime.Stream, args...)
	return err
}
