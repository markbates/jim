package seed

import (
	"context"
	"fmt"
)

// Users puts all of the users into all of the databases
func Users(ctx context.Context, args []string) error {
	fmt.Println("loading users", args)
	return nil
}

func NotATask() {}
