package main

import (
	"log"
	"os"

	"github.com/dnncomp/telega/app/cmd"
	"github.com/umputun/go-flags"
)

// Opts with all cli commands and flags
type Opts struct {
	TokenCmd    cmd.TokenCommand    `command:"token" description:"get token"`
	BotInfoCmd  cmd.BotInfoCommand  `command:"botinfo" description:"get bot info"`
	ChatInfoCmd cmd.ChatInfoCommand `command:"chatinfo" description:"get chat info"`

	URL   string `long:"url" env:"TELEGA_URL" required:"true" default:"https://api.telegram.org" description:"telegram api URL"`
	Token string `short:"t" long:"token" env:"TELEGA_BOT_TOKEN" required:"true" default:"XXXXXX" description:"bot token"`
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
