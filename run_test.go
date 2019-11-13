package jim

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Run(t *testing.T) {
	r := require.New(t)

	in := strings.NewReader("")
	out := &bytes.Buffer{}
	ew := &bytes.Buffer{}

	ctx := tContext{
		Context: context.Background(),
		stdin:   in,
		stdout:  out,
		stderr:  ew,
	}

	k, err := New([]string{"Jim"})
	r.NoError(err)

	err = Run(ctx, k)
	r.NoError(err)

	exp := strings.TrimSpace(jim)
	act := strings.TrimSpace(out.String())
	r.Equal(exp, act)

	k, err = New([]string{"unknown:BadTask"})
	r.NoError(err)

	err = Run(ctx, k)
	r.Error(err)
}
