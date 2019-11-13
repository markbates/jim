package jim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Help(t *testing.T) {
	r := require.New(t)

	k, err := New([]string{"Jim"})
	r.NoError(err)

	s, err := Help(k)
	r.NoError(err)

	exp := "func Jim(ctx context.Context, args []string) error"
	r.Contains(s, exp)
}
