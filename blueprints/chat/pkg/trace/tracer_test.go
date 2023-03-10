package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run(`ensure that trace will be correct initialized`, func(t *testing.T) {
		var buf bytes.Buffer
		tracer := New(&buf)
		if tracer == nil {
			t.Error("Return from New should not be nil")
		} else {
			tracer.Trace("Hello trace package.")
			if buf.String() != "Hello trace package.\n" {
				t.Errorf("Trace should not write '%s'.", buf.String())
			}
		}
	})
}

func TestOff(t *testing.T) {
	t.Run(`ensure tracing doesn't panic when trace fails to start`, func(t *testing.T) {
		var silentTracer Tracer = Off()
		silentTracer.Trace("something")
	})
}
