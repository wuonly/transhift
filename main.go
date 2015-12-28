package main

import (
    "github.com/transhift/transhift/transhift"
    "github.com/codegangsta/cli"
    "os"
)

func main() {
    app := cli.NewApp()

    app.Name = "Transhift"
    app.Usage = "Peer-to-peer file transfers"
    app.Version = "0.1.0"

    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name: "app-dir",
            Value: "",
            Usage: "application directory",
        },
    }

    app.Commands = []cli.Command{
        {
            Name: "download",
            Aliases: []string{"dl"},
            Usage: "download from a peer",
            Action: transhift.Download,
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name: "destination, d",
                    Value: "",
                    Usage: "destination file",
                },
            },
        },
        {
            Name: "upload",
            Aliases: []string{"ul"},
            Usage: "Upload to a peer",
            ArgsUsage: "<peer> <file>",
            Action: transhift.Upload,
        },
    }

    app.Run(os.Args)
}
