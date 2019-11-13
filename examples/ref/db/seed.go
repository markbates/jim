package db

import (
	"context"
	"fmt"
)

func Seed(ctx context.Context, args []string) error {
	fmt.Println("seeding the db ", args)
	return nil
}
