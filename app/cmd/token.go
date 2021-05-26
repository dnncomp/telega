package cmd

import (
	"fmt"
)

type TokenCommand struct {
	CommonOpts
}

func (c *TokenCommand) Execute(_ []string) error {

	fmt.Printf("token: %+v", c.Token)
	return nil
}
