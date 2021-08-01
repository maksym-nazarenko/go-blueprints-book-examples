package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Tracer returned from New() is nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("\nexpected: %v\ngot: %v\n", "Hello trace package.", buf.String())
		}
	}
}
