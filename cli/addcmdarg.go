package cli

func addcmdarg() {
	Auth = Root.PersistentFlags().String("auth", "none", "认证信息")
	Host = Root.PersistentFlags().String("host", "", "监听/连接地址 (default \"ALL\")")
	Port = Root.PersistentFlags().StringP("port", "p", "4870", "监听/连接端口")
}
