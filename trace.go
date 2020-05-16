package gerr

import (
	"runtime"
	"strconv"
)

type traceCollection struct {
	logs []traceLog
}

type traceLog struct {
	pc   uintptr
	file string
	line int
}

// it returns trace skip how many of callers
func trace(skip int) (tc traceCollection) {
	var next = true
	for i := skip + 1; next; i++ {
		var log traceLog
		if log.pc, log.file, log.line, next = runtime.Caller(i); next {
			tc.logs = append(tc.logs, log)
		}
	}
	return
}

func (tc traceCollection) log(padding string) (msg string) {
	for _, log := range tc.logs {
		msg += padding + "at " + runtime.FuncForPC(log.pc).Name() + " (" + log.file + ":" + strconv.Itoa(log.line) + ")\n"
	}
	return
}

func (tc *traceCollection) clean() {
	tc.logs = nil
}
