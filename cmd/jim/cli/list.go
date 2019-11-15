package cli

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/markbates/jim"
)

func List(ctx context.Context, args []string) error {
	opts := struct {
		*flag.FlagSet
		json bool
	}{
		FlagSet: flag.NewFlagSet("jim list", flag.ExitOnError),
	}
	opts.BoolVar(&opts.json, "json", false, "display as json")

	if err := opts.Parse(args); err != nil {
		return err
	}

	args = opts.Args()

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	tasks, err := jim.List(pwd)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		return nil
	}

	stdout := jim.Stdout(ctx)

	if opts.json {
		enc := json.NewEncoder(stdout)
		enc.SetIndent("", " ")
		return enc.Encode(tasks)
	}

	w := tabwriter.NewWriter(stdout, 0, 0, 0, ' ', tabwriter.Debug)
	defer w.Flush()
	fmt.Fprintln(w, "Task \t GoDoc")
	fmt.Fprintln(w, "---- \t -----")
	for _, t := range tasks {
		fmt.Fprintf(w,
			"%s \t %s\n",
			t.String(),
			t.GoDoc(),
		)
	}
	return nil
}
