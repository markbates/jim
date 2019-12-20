package jim

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type errNoTask struct {
	t   *Task
	err error
}

func (e errNoTask) Error() string {
	return e.Error()
}

func (e errNoTask) Task() *Task {
	return e.t
}

func Run(ctx context.Context, t *Task) error {
	c := exec.CommandContext(ctx, "go", "doc", fmt.Sprintf("%s.%s", t.Pkg, t.Name))
	if err := c.Run(); err != nil {
		return errNoTask{
			err: err,
			t:   t,
		}
	}

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

	c = command(ctx, "go", args...)
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}

const tmpl = `
package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"{{.Pkg}}"
)

func main() {
	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	if err := {{.Sel}}.{{.Name}}(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
`
