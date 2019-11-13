package jim

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Jim(t *testing.T) {
	r := require.New(t)

	bb := &bytes.Buffer{}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "stdout", bb)

	r.NoError(Jim(ctx, []string{}))

	exp := strings.TrimSpace(jim)
	act := strings.TrimSpace(bb.String())
	r.Contains(exp, act)
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
