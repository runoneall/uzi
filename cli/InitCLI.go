package cli

import "github.com/spf13/cobra"

func InitCLI(use string) {
	once.Do(func() {
		Root = &cobra.Command{
			Use:   use,
			Short: "为提供内网 ip 的 vps 编写的终端通讯系统",
			Run: func(cmd *cobra.Command, args []string) {
				Root.Help()
			},
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		}

		addsubcmd()
		addcmdarg()
	})
}
