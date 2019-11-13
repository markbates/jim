package task

import (
	"context"
	"fmt"
)

// Something does something
func Something(ctx context.Context, args []string) error {
	fmt.Println("doing something", args)
	return nil
}

// Another... take a guess. :)
func Another(ctx context.Context, args []string) error {
	fmt.Println("doing another thing", args)
	return nil
}

func NotATask(int, int) error {
	return nil
}
