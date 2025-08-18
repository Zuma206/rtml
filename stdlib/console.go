package stdlib

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/zuma206/rtml/runtime"
)

func openConsole(rt *runtime.Runtime) {
	console := rt.VM.NewObject()
	console.Set("log", newLog(rt))
	rt.VM.Set("console", console)
}

func newLog(rt *runtime.Runtime) any {
	return func(values ...goja.Value) {
		for i, value := range values {
			if i != 0 {
				fmt.Fprint(rt.Log, " ")
			}
			fmt.Fprint(rt.Log, value.String())
		}
		fmt.Fprintln(rt.Log)
	}
}
