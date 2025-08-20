package stdlib

import (
	"github.com/zuma206/rtml/runtime"
)

func OpenStdlib(rt *runtime.Runtime) {
	openConsole(rt)
	openTempl(rt)
}
