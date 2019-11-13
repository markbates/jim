package jim

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
)

func Test_Stdin(t *testing.T) {
	ctx := context.Background()

	r := Stdin(ctx)
	if r != os.Stdin {
		t.Fatalf("expected os.Stdin got %s", r)
	}

	var in io.Reader
	in = &bytes.Buffer{}
	ctx = context.WithValue(ctx, "stdin", in)

	r = Stdin(ctx)
	if r != in {
		t.Fatalf("expected %T got %s", in, r)
	}

	in = &bytes.Buffer{}
	ctx = tContext{
		Context: ctx,
		stdin:   in,
	}

	r = Stdin(ctx)
	if r != in {
		t.Fatalf("expected %T got %s", in, r)
	}
}

func Test_Stdout(t *testing.T) {
	ctx := context.Background()

	w := Stdout(ctx)
	if w != os.Stdout {
		t.Fatalf("expected os.Stdout got %s", w)
	}

	var out io.Writer
	out = &bytes.Buffer{}
	ctx = context.WithValue(ctx, "stdout", out)

	w = Stdout(ctx)
	if w != out {
		t.Fatalf("expected %T got %s", out, w)
	}

	out = &bytes.Buffer{}
	ctx = tContext{
		Context: ctx,
		stdout:  out,
	}

	w = Stdout(ctx)
	if w != out {
		t.Fatalf("expected %T got %s", out, w)
	}
}

func Test_Stderr(t *testing.T) {
	ctx := context.Background()

	w := Stderr(ctx)
	if w != os.Stderr {
		t.Fatalf("expected os.Stderr got %s", w)
	}

	var out io.Writer
	out = &bytes.Buffer{}
	ctx = context.WithValue(ctx, "stderr", out)

	w = Stderr(ctx)
	if w != out {
		t.Fatalf("expected %T got %s", out, w)
	}

	out = &bytes.Buffer{}
	ctx = tContext{
		Context: ctx,
		stderr:  out,
	}

	w = Stderr(ctx)
	if w != out {
		t.Fatalf("expected %T got %s", out, w)
	}
}

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
