package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/markbates/jim"
)

type cmdOptions struct {
	*flag.FlagSet
	help bool
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	oargs := os.Args[1:]

	opts := cmdOptions{
		FlagSet: flag.NewFlagSet("", flag.ContinueOnError),
	}
	opts.BoolVar(&opts.help, "help", false, "display help")
	opts.BoolVar(&opts.help, "h", false, "display help")

	if err := opts.Parse(oargs); err != nil {
		return err
	}

	oargs = opts.Args()
	if opts.help && len(oargs) == 0 {
		opts.Usage()
		return nil
	}

	t, err := jim.New(oargs)
	if err != nil {
		return err
	}

	if opts.help {
		return taskHelp(t)
	}

	return runTask(t)
}

func taskHelp(t *jim.Task) error {
	c := exec.Command("go", "doc", fmt.Sprintf("%s.%s", t.Pkg, t.Name))
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

func runTask(t *jim.Task) error {
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

	fmt.Fprintf(f, tmpl, t.Pkg, t.Sel, t.Name)

	args := []string{"run", out}
	args = append(args, t.Args...)

	c := exec.Command("go", args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

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
)
import "%s"

func main() {
		ctx := context.Background()
		if err := %s.%s(ctx, os.Args[1:]); err != nil {
				log.Fatal(err)
		}
}
`
