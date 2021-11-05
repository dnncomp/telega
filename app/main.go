package main

import (
	"github.com/dnncomp/telega/app/logger"
	"os"

	"github.com/dnncomp/telega/app/cmd"
	"github.com/umputun/go-flags"
)

// Opts with all cli commands and flags
type Opts struct {
	TokenCmd    cmd.TokenCommand    `command:"token" description:"- get token" hidden:"true"`
	BotInfoCmd  cmd.BotInfoCommand  `command:"botinfo" description:"- get bot info"`
	ChatInfoCmd cmd.ChatInfoCommand `command:"chatinfo" description:"- get chat info"`
	SendCmd     cmd.SendCommand     `command:"send" description:"- send message to chat"`

	URL   string `long:"url" env:"TELEGA_URL" required:"true" default:"https://api.telegram.org" description:"telegram api URL"`
	Token string `short:"t" long:"token" env:"TELEGA_BOT_TOKEN" required:"true" description:"bot token"`

	LogerGroup struct {
		Path       string `long:"path" env:"PATH" default:"telegalog/" description:"папка хранения логов"`
		Level      string `long:"level" env:"LEVEL" default:"WARN" description:"уровень логирования"`
		MaxSize    int    `long:"size" env:"SIZE" default:"10" description:"макс. размер файла лога (MB)"`
		MaxBackups int    `long:"backups" env:"BACKUPS" default:"5" description:"макс. количество файлов лога"`
		MaxAge     int    `long:"age" env:"AGE" default:"30" description:"макс. возраст лога (дней)"`
	} `group:"log" namespace:"log" env-namespace:"LOG"`

	Dbg bool `short:"d" long:"dbg" env:"DEBUG" description:"включен режим отладки"`
}

func main() {

	var opt Opts
	p := flags.NewParser(&opt, flags.Default)
	p.CommandHandler = func(command flags.Commander, args []string) error {

		// Init logger
		if opt.Dbg {
			opt.LogerGroup.Level = "DEBUG"
		}

		log := logger.New(opt.LogerGroup.Path, opt.LogerGroup.Level,
			opt.LogerGroup.MaxSize, opt.LogerGroup.MaxBackups, opt.LogerGroup.MaxAge)

		c := command.(cmd.CommonOptionsCommander)
		c.SetCommon(cmd.CommonOpts{
			URL:   opt.URL,
			Token: opt.Token,
			Log:   log,
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
