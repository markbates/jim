package jim

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_List(t *testing.T) {
	r := require.New(t)

	pwd, err := os.Getwd()
	r.NoError(err)

	root := filepath.Join(pwd, "examples", "ref")

	tasks, err := List(root)
	r.NoError(err)
	r.Len(tasks, 4)

	var act []string
	for _, t := range tasks {
		act = append(act, t.String())
	}

	exp := []string{
		"ref/db/seed:Users",
		"ref/db:Seed",
		"ref/task:Another",
		"ref/task:Something",
	}
	r.Equal(exp, act)
}
