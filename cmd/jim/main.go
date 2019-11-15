package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/markbates/jim"
	"github.com/markbates/jim/cmd/jim/cli"
)

type cmdOptions struct {
	*flag.FlagSet
	help bool
	json bool
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	oargs := os.Args[1:]

	opts := cmdOptions{
		FlagSet: flag.NewFlagSet("jim", flag.ContinueOnError),
	}
	opts.BoolVar(&opts.help, "h", false, "display help")

	if err := opts.Parse(oargs); err != nil {
		return err
	}

	oargs = opts.Args()
	if opts.help && len(oargs) == 0 {
		opts.Usage()
		return nil
	}

	if len(oargs) > 0 && oargs[0] == "list" {
		ctx := context.Background()
		return cli.List(ctx, oargs[1:])
	}

	t, err := jim.New(oargs)
	if err != nil {
		return err
	}

	if opts.help {
		return taskHelp(t)
	}

	ctx := context.Background()
	return jim.Run(ctx, t)
}

func taskHelp(t *jim.Task) error {
	s, err := jim.Help(t)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
