package jim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	_, err := New([]string{})
	r.Error(err)

	_, err = New([]string{"x"})
	r.Error(err)

	k, err := New([]string{"db/seed.Users", "a", "b", "c"})
	r.NoError(err)
	r.NotZero(k)
	r.Equal("github.com/markbates/jim/db/seed", k.Pkg)
	r.Equal("seed", k.Sel)
	r.Equal("Users", k.Name)
	r.Equal([]string{"a", "b", "c"}, k.Args)

	k, err = New([]string{"Users"})
	r.NoError(err)

	r.NotZero(k)
	r.Equal("github.com/markbates/jim", k.Pkg)
	r.Equal("jim", k.Sel)
	r.Equal("Users", k.Name)

	k, err = New([]string{"github.com/markbates/jim.Jim"})
	r.NoError(err)

	r.NotZero(k)
	r.Equal("github.com/markbates/jim", k.Pkg)
	r.Equal("jim", k.Sel)
	r.Equal("Jim", k.Name)
}
