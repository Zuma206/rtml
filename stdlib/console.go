package stdlib

import (
	"github.com/dop251/goja"
)

func openConsole(vm *goja.Runtime) {
	console := vm.NewObject()
	console.Set("log", log)
	vm.Set("console", console)
}

func log(values ...goja.Value) {
	for i, value := range values {
		if i != 0 {
			print(" ")
		}
		print(value.String())
	}
	println()
}
