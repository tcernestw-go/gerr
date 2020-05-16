package gerr

import "testing"

func Test_traceCollection_log(t *testing.T) {
	tests := []struct {
		name    string
		tc      traceCollection
		padding string
		wantMsg string
	}{
		{
			name:    "Normal Trace",
			tc:      trace(0),
			padding: " ",
			wantMsg: " at gerr.Test_traceCollection_log (/media/prime/InternalHDD/own/project/lib/go/src/gerr/trace_test.go:14)\n at testing.tRunner (/usr/local/go/src/testing/testing.go:909)\n at runtime.goexit (/usr/local/go/src/runtime/asm_amd64.s:1357)\n",
		},
		{
			name:    "Negative Skip",
			tc:      trace(-1),
			padding: " ",
			wantMsg: " at gerr.trace (/media/prime/InternalHDD/own/project/lib/go/src/gerr/trace.go:23)\n at gerr.Test_traceCollection_log (/media/prime/InternalHDD/own/project/lib/go/src/gerr/trace_test.go:20)\n at testing.tRunner (/usr/local/go/src/testing/testing.go:909)\n at runtime.goexit (/usr/local/go/src/runtime/asm_amd64.s:1357)\n",
		},
		{
			name:    "Invalid Negative Skip",
			tc:      trace(-3),
			padding: " ",
			wantMsg: " at runtime.Caller (/usr/local/go/src/runtime/extern.go:182)\n at runtime.Caller (/usr/local/go/src/runtime/extern.go:182)\n at gerr.trace (/media/prime/InternalHDD/own/project/lib/go/src/gerr/trace.go:23)\n at gerr.Test_traceCollection_log (/media/prime/InternalHDD/own/project/lib/go/src/gerr/trace_test.go:26)\n at testing.tRunner (/usr/local/go/src/testing/testing.go:909)\n at runtime.goexit (/usr/local/go/src/runtime/asm_amd64.s:1357)\n",
		},
		{
			name:    "Positive Skip",
			tc:      trace(1),
			padding: " ",
			wantMsg: " at testing.tRunner (/usr/local/go/src/testing/testing.go:909)\n at runtime.goexit (/usr/local/go/src/runtime/asm_amd64.s:1357)\n",
		},
		{
			name:    "Invalid Positive Skip",
			tc:      trace(3),
			padding: " ",
			wantMsg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMsg := tt.tc.log(tt.padding); gotMsg != tt.wantMsg {
				t.Errorf("traceCollection.log() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
