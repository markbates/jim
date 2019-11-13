package jim

import (
	"context"
	"io"
	"os"
)

type Context struct {
	context.Context
	in  io.Reader
	out io.Writer
	err io.Writer
}

func (c Context) Stdin() io.Reader {
	if r, ok := c.Value("stdin").(io.Reader); ok {
		return r
	}

	if c.in != nil {
		return c.in
	}

	return os.Stdin
}

func (c Context) Stdout() io.Writer {
	if r, ok := c.Value("stdout").(io.Writer); ok {
		return r
	}

	if c.out != nil {
		return c.out
	}

	return os.Stdout
}

func (c Context) Stderr() io.Writer {
	if r, ok := c.Value("stderr").(io.Writer); ok {
		return r
	}

	if c.out != nil {
		return c.out
	}

	return os.Stderr
}

func NewContext(ctx context.Context) Context {
	return Context{
		Context: ctx,
		in:      os.Stdin,
		out:     os.Stdout,
		err:     os.Stderr,
	}
}
