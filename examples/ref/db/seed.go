package db

import (
	"context"
	"fmt"
)

func Seed(ctx context.Context, args []string) error {
	fmt.Println("seeding the db ", args)
	return nil
}

func NotATask(ctx context.Context, args []int) error {
	return nil
}
