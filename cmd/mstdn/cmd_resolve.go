package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mattn/go-mastodon"
	"github.com/urfave/cli/v2"
)

func cmdResolve(c *cli.Context) error {
	client := c.App.Metadata["client"].(*mastodon.Client)
	if !c.Args().Present() {
		return errors.New("arguments required")
	}
	for i := 0; i < c.NArg(); i++ {
		results, err := client.Search(context.Background(), c.Args().Get(i), true)
		if err != nil {
			return err
		}
		if len(results.Statuses) == 0 {
			return errors.New("unable to resolve link")
		}
		for _, s := range results.Statuses {
			fmt.Println(s.ID)
		}
	}
	return nil
}