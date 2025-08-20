package stdlib

import (
	"github.com/dop251/goja"
	"github.com/zuma206/rtml/runtime"
)

func openTempl(rt *runtime.Runtime) {
	target := rt.VM.NewObject()
	rt.VM.Set("templ", rt.VM.NewProxy(target, &goja.ProxyTrapConfig{
		Get: func(_ *goja.Object, property string, _ goja.Value) (value goja.Value) {
			template, ok := rt.Templates[property]
			if !ok {
				return goja.Undefined()
			}
			return rt.VM.ToValue(func() {
				for child := range template.ChildNodes() {
					if err := rt.Eval(child); err != nil {
						panic(rt.VM.NewGoError(err))
					}
				}
			})
		},
	}))
}
