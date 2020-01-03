package jim

import (
	"context"
	"os/exec"
)

func command(ctx context.Context, n string, args ...string) *exec.Cmd {
	c := exec.CommandContext(ctx, n, args...)
	// fmt.Println(c.Args)
	c.Stdin = Stdin(ctx)
	c.Stdout = Stdout(ctx)
	c.Stderr = Stderr(ctx)
	return c
}
