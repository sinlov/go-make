package subcommand_init

import (
	"github.com/sinlov/go-make/command"
	"github.com/sinlov/go-make/internal/d_log"
	"github.com/sinlov/go-make/internal/urfave_cli"
	"github.com/urfave/cli/v2"
)

const commandName = "init"

var commandEntry *InitCommand

type InitCommand struct {
	isDebug bool
}

func (n *InitCommand) Exec() error {
	d_log.Debugf("-> Exec subCommand [ %s ]", commandName)

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "lib",
			Usage: "Use a library template",
			Value: false,
		},
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "Set the resulting package name, defaults to the directory name",
		},
	}
}

func withEntry(c *cli.Context) (*InitCommand, error) {
	d_log.Debugf("-> withEntry subCommand [ %s ]", commandName)

	if c.Bool("lib") {
		d_log.Info("new lib mode")
	}
	globalEntry := command.CmdGlobalEntry()
	return &InitCommand{
		isDebug: globalEntry.Verbose,
	}, nil
}

func action(c *cli.Context) error {
	d_log.Debugf("-> Sub Command action [ %s ] start", commandName)
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	commandEntry = entry
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "init config",
			Action: action,

			Flags: flag(),
			//Flags: urfave_cli.UrfaveCliAppendCliFlag(flag(), constant.PlatformFlags()),
		},
	}
}
