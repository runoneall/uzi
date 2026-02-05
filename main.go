package main

import (
	"os"
	"path/filepath"
	"uzi/cli"
	"uzi/client"
	"uzi/server"

	"github.com/spf13/cobra"
)

func main() {
	selfname, err := os.Executable()
	if err != nil {
		panic(err)
	}

	cli.InitCLI(filepath.Base(selfname))
	cli.Serve.Run = func(cmd *cobra.Command, args []string) { server.Serve() }
	cli.Connect.Run = func(cmd *cobra.Command, args []string) { client.Connect() }
	cli.Execute()
}
