package jim

import (
	"context"
	"os"
	"path/filepath"
	"text/template"
)

func Run(ctx context.Context, t *Task) error {
	od := filepath.Join(t.Dir, ".jim")
	out := filepath.Join(od, "main.go")
	os.MkdirAll(od, 0755)
	defer os.RemoveAll(od)
	defer func() {
		if err := recover(); err != nil {
			os.RemoveAll(od)
		}
	}()

	f, err := os.Create(out)
	if err != nil {
		return err
	}
	defer f.Close()

	mpl, err := template.New("main.go").Parse(tmpl)
	if err != nil {
		return err
	}
	if err := mpl.Execute(f, t); err != nil {
		return err
	}

	args := []string{"run", out}
	args = append(args, t.Args...)

	c := command(ctx, "go", args...)
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}

const tmpl = `package main

import(
		"context"
		"log"
		"os"
		"{{.Pkg}}"
)

func main() {
		ctx := context.Background()
		if err := {{.Sel}}.{{.Name}}(ctx, os.Args[1:]); err != nil {
				log.Fatal(err)
		}
}
`
