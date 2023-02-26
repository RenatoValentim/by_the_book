package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...any)
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...any) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...any) {}

func Off() Tracer {
	return &nilTracer{}
}
