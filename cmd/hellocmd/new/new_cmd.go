package new

import (
	"github.com/urfave/cli"
)

var Cmd = cli.Command{
	Name:            "new",
	Aliases:         []string{"n"},
	Usage:           "Create hellocmd template project",
	Action:          CreateProject,
	SkipFlagParsing: false,
	UsageText:       ProjectHelpTemplate,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "dir, d",
			Usage:       "Specify the directory of the project",
			Destination: &project.Path,
		},
	},
}
