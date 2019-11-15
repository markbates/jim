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
	Args []string
	Pkg  string // github.com/foo/bar
	Sel  string // bar
	Name string // Something
}

func (t Task) String() string {
	return fmt.Sprintf("%s.%s", t.Pkg, t.Name)
}

func (t Task) GoDoc() string {
	return fmt.Sprintf("https://godoc.org/%s#%s", t.Pkg, t.Name)
}

func New(args []string) (*Task, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("missing task name")
	}

	info, err := here.Current()
	if err != nil {
		return nil, err
	}

	t := &Task{
		Info: info,
		Args: args[1:],
	}

	ext := filepath.Ext(args[0])
	t.Pkg = strings.TrimSuffix(args[0], ext)
	t.Name = strings.TrimPrefix(ext, ".")

	if len(ext) == 0 {
		t.Name = args[0]
		t.Pkg = info.ImportPath
	}

	if !strings.HasPrefix(t.Pkg, info.ImportPath) {
		t.Pkg = path.Join(info.ImportPath, t.Pkg)
	}

	t.Sel = path.Base(t.Pkg)

	rn, _ := utf8.DecodeRuneInString(t.Name)
	if !unicode.IsUpper(rn) {
		return nil, fmt.Errorf("functions must be publicly exported %s", args[0])
	}
	return t, nil
}
