package cmd

import (
	"fmt"
)

type GetMeCommand struct {
	CommonOpts
}

func (c *GetMeCommand) Execute(_ []string) error {
	fmt.Printf("getme: %+v", c.Token)
	return nil
}
