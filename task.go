package jim

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gobuffalo/here"
)

type Task struct {
	here.Info
	BuildArgs []string // -tags foo -v
	Args      []string
	Pkg       string // github.com/foo/bar
	Sel       string // bar
	Name      string // Something
}

func (t Task) String() string {
	return fmt.Sprintf("%s:%s", t.Pkg, t.Name)
}

func New(args []string) (*Task, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("missing task name")
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

	in := args[0]
	t.Name = filepath.Ext(in)
	if len(t.Name) == 0 {
		t.Name = in
	}
	t.Pkg = strings.TrimSuffix(in, t.Name)
	t.Name = strings.TrimPrefix(t.Name, ".")

	if !strings.HasPrefix(t.Pkg, info.Module.Path) {
		t.Pkg = path.Join(info.Module.Path, t.Pkg)
	}

	rn, _ := utf8.DecodeRuneInString(t.Name)
	if !unicode.IsUpper(rn) {
		return nil, fmt.Errorf("functions must be exported and start with an upper case: %s", t.Name)
	}

	t.Sel = path.Base(t.Pkg)
	return t, nil
}
