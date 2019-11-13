package jim

import (
	"fmt"
	"path"
	"strings"
	"unicode"
	"unicode/utf8"

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
	if len(parts) < 1 {
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

	var pkg string
	var name string

	if len(parts) == 1 {
		pkg = info.ImportPath
		name = parts[0]
	} else {
		pkg = strings.Join(parts[:len(parts)-1], "/")
		pkg = path.Join(info.ImportPath, pkg)
		name = parts[len(parts)-1]
	}
	t.Pkg = pkg
	t.Name = name

	rn, _ := utf8.DecodeRuneInString(t.Name)
	if !unicode.IsUpper(rn) {
		return nil, fmt.Errorf("functions must be exported and start with an upper case: %s", t.Name)
	}

	t.Sel = path.Base(t.Pkg)
	return t, nil
}
