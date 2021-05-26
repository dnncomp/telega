package main

import (
	"github.com/dnncomp/telega/app/cmd"
	"github.com/umputun/go-flags"
	"log"
	"os"
)

// Opts with all cli commands and flags
type Opts struct {
	TokenCmd cmd.TokenCommand `command:"token" description:"get token"`

	URL   string `long:"url" env:"URL" required:"true" default:"https://api.telegram.org" description:"telegram api URL"`
	Token string `short:"t" long:"token" env:"TOKEN" required:"true" default:"XXXXXX" description:"bot token"`
}

func main() {

	var opt Opts
	p := flags.NewParser(&opt, flags.Default)
	p.CommandHandler = func(command flags.Commander, args []string) error {

		c := command.(cmd.CommonOptionsCommander)
		c.SetCommon(cmd.CommonOpts{
			URL:   opt.URL,
			Token: opt.Token,
		})

		err := c.Execute(args)
		if err != nil {
			log.Printf("[ERROR] failed with %+v", err)
		}
		return nil
	}

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
