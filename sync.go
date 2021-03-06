package main

import (
	"context"

	"github.com/urfave/cli"
)

func Sync(c *cli.Context) error {
	client := GetClient(c)

	err := client.Sync(context.Background())
	if err != nil {
		return err
	}
	err = WriteCache(default_cache_path, client.Store)
	if err != nil {
		return err
	}

	return nil
}
