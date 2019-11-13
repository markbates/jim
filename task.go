package jim

import (
	"fmt"
	"path"
	"strings"

	"github.com/gobuffalo/here"
)

type Task struct {
	here.Info
	Args []string
	Pkg  string // github.com/foo/bar
	Sel  string // bar
	Name string // Something
}

func New(args []string) (*Task, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("missing task name")
	}

	parts := strings.Split(args[0], ":")
	if len(parts) < 2 {
		return nil, fmt.Errorf("malformed task name %s", args[0])
	}

	h := here.New()

	info, err := h.Current()
	if err != nil {
		return nil, err
	}

	t := &Task{
		Info: info,
		Args: args[1:],
	}

	pkgs := parts[:len(parts)-1]
	t.Sel = pkgs[len(pkgs)-1]
	t.Name = parts[len(parts)-1]

	t.Pkg = path.Join(pkgs...)
	t.Pkg = path.Join(info.ImportPath, t.Pkg)
	return t, nil
}
