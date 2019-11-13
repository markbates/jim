package jim

import (
	"context"
	"io"
	"os"
)

// Stdin retuns an io.Reader representing the stdin.
// If the ctx implements Iny that is returned.
// If the ctx.Value("stdin") is present, and an io.Reader,
// that is returned.
// If no other io.Reader is found, then os.Stdin is returned.
func Stdin(ctx context.Context) io.Reader {
	if r, ok := ctx.(Iny); ok {
		return r.Stdin()
	}

	if r, ok := ctx.Value("stdin").(io.Reader); ok {
		return r
	}

	return os.Stdin
}

// Stdout retuns an io.Writer representing the stdout.
// If the ctx implements Outy that is returned.
// If the ctx.Value("stdout") is present, and an io.Writer,
// that is returned.
// If no other io.Writer is found, then os.Stdout is returned.
func Stdout(ctx context.Context) io.Writer {
	if r, ok := ctx.(Outy); ok {
		return r.Stdout()
	}

	if r, ok := ctx.Value("stdout").(io.Writer); ok {
		return r
	}

	return os.Stdout
}

// Stderr retuns an io.Writer representing the stderr.
// If the ctx implements Erry that is returned.
// If the ctx.Value("stderr") is present, and an io.Writer,
// that is returned.
// If no other io.Writer is found, then os.Stderr is returned.
func Stderr(ctx context.Context) io.Writer {
	if r, ok := ctx.(Erry); ok {
		return r.Stderr()
	}

	if r, ok := ctx.Value("stderr").(io.Writer); ok {
		return r
	}

	return os.Stderr
}
