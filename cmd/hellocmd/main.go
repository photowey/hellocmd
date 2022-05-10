package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"

	"github.com/photowey/hellocmd/cmd/hellocmd/new"
)

const (
	Version = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "hellocmd"
	app.Usage = "hellocmd cmder tools"
	app.Version = Version
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "photowey",
			Email: "photowey@gmail.com",
		},
	}
	app.Copyright = "(c) 2022 photowey <photowey@gmail.com>"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello hellocmd!")
		return nil
	}

	app.Commands = []cli.Command{
		new.Cmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
