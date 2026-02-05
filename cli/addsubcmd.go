package cli

import (
	"github.com/spf13/cobra"
)

func addsubcmd() {
	Serve = &cobra.Command{
		Use:   "serve",
		Short: "启动后端服务",
	}

	Connect = &cobra.Command{
		Use:   "connect",
		Short: "连接后端服务",
	}

	Root.AddCommand(
		Serve,
		Connect,
	)
}
