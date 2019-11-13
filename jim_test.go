package jim

import (
	"context"
	"io"
)

type tContext struct {
	context.Context
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

func (t tContext) Stdin() io.Reader {
	return t.stdin
}

func (t tContext) Stdout() io.Writer {
	return t.stdout
}

func (t tContext) Stderr() io.Writer {
	return t.stderr
}
