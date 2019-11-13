package jim

import (
	"context"
	"io"
)

func In(ctx context.Context) io.Reader {
	c := NewContext(ctx)
	return c.Stdin()
}

func Out(ctx context.Context) io.Writer {
	c := NewContext(ctx)
	return c.Stdout()
}

func Err(ctx context.Context) io.Writer {
	c := NewContext(ctx)
	return c.Stderr()
}
