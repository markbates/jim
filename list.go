package jim

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gobuffalo/here"
)

var defaultIgnoredFolders = []string{".", "_", "vendor", "node_modules", "testdata"}

// List parses the AST at, and below, the given root looking for Jim tasks.
func List(root string) ([]*Task, error) {
	var tasks []*Task
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}

		base := filepath.Base(path)
		for _, x := range defaultIgnoredFolders {
			if strings.HasPrefix(base, x) {
				return filepath.SkipDir
			}
		}

		tsk, err := parse(path)
		if err != nil {
			return err
		}
		tasks = append(tasks, tsk...)

		return nil
	})

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].String() < tasks[j].String()
	})

	return tasks, err
}

func parse(root string) ([]*Task, error) {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, root, nil, 0)
	if err != nil {
		return nil, err
	}

	h := here.New()
	info, err := h.Dir(root)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	for _, pkg := range pkgs {
		for _, pf := range pkg.Files {
			for _, d := range pf.Decls {
				k, err := fromDecl(info, d)
				if err != nil {
					return nil, err
				}
				if k != nil {
					tasks = append(tasks, k)
				}
			}
		}
	}
	return tasks, nil
}

func fromDecl(info here.Info, d ast.Node) (*Task, error) {
	fn, ok := d.(*ast.FuncDecl)
	if !ok {
		return nil, nil
	}

	name := fn.Name.Name
	rn, _ := utf8.DecodeRuneInString(name)
	if !unicode.IsUpper(rn) {
		return nil, nil
	}

	ft := fn.Type
	if ft.Params == nil {
		return nil, nil
	}
	parms := ft.Params.List
	if len(parms) < 2 {
		return nil, nil
	}

	sel, ok := parms[0].Type.(*ast.SelectorExpr)
	if !ok {
		return nil, nil
	}

	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return nil, nil
	}

	arg1 := fmt.Sprintf("%s.%s", id.Name, sel.Sel.Name)
	if arg1 != "context.Context" {
		return nil, nil
	}

	art, ok := parms[1].Type.(*ast.ArrayType)
	if !ok {
		return nil, nil
	}

	id, ok = art.Elt.(*ast.Ident)
	if !ok {
		return nil, nil
	}

	if id.Name != "string" {
		return nil, nil
	}

	return &Task{
		Name: name,
		Info: info,
		Pkg:  info.ImportPath,
		Sel:  info.Name,
	}, nil

}
