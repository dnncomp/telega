package cmd

import (
	"log"
	"strings"
)

var Log *log.Logger

// CommonOptionsCommander extends flags.Commander with SetCommon
// All commands should implement this interfaces
type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}

// CommonOpts sets externally from main, shared across all commands
type CommonOpts struct {
	URL   string
	Token string
	Log   *log.Logger
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (c *CommonOpts) SetCommon(commonOpts CommonOpts) {
	c.URL = strings.TrimSuffix(commonOpts.URL, "/") // allow URL with trailing /
	c.Token = commonOpts.Token
	c.Log = commonOpts.Log
}
