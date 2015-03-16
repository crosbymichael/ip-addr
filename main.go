package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ip-addr"
	app.Usage = "return the IP address for a given interface name"
	app.Version = "1"
	app.Author = "@crosbymichael"
	app.Email = "crosbymichael@gmail.com"
	app.Action = func(context *cli.Context) {
		i, err := net.InterfaceByName(context.Args().First())
		if err != nil {
			logrus.Fatal(err)
		}
		addrs, err := i.Addrs()
		if err != nil {
			logrus.Fatal(err)
		}
		// get the highest order ip
		if len(addrs) == 0 {
			return
		}
		for _, a := range addrs {
			n, ok := a.(*net.IPNet)
			if !ok {
				continue
			}
			fmt.Print(n.IP)
			return
		}
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
