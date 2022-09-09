package impl

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
)

func GetFuncName() string {
	var buffer bytes.Buffer
	pc := make([]uintptr, 10)
	runtime.Callers(4, pc)
	frame, _ := runtime.CallersFrames(pc).Next()
	function := frame.Function
	line := frame.Line
	buffer.WriteString(function)
	buffer.WriteString(fmt.Sprintf(":%d", line))

	return filepath.Base(buffer.String())
}
