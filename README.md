# uzi

为提供内网 ip 的 vps 编写的终端通讯系统

## 使用方法

首先到 https://github.com/runoneall/uzi/releases 下载最新二进制文件

```plaintext
$ uzi --help
为提供内网 ip 的 vps 编写的终端通讯系统

Usage:
  uzi [flags]
  uzi [command]

Available Commands:
  connect     连接后端服务
  help        Help about any command
  serve       启动后端服务

Flags:
      --auth string   认证信息 (default "none")
  -h, --help          help for uzi
      --host string   监听/连接地址 (default "ALL")
  -p, --port string   监听/连接端口 (default "4870")

Use "uzi [command] --help" for more information about a command.
```

## 扩展包使用

uzi 提供了一些扩展包，对实时通讯的一些功能进行了封装，可用于集成到你的项目中

- `uzi/conn` 该包提供了对大量网络连接 `net.Conn` 的 `添加` `删除` `广播` 功能的支持，并进行了算法优化，时间复杂度 `O(1)`
- `uzi/history` 该包提供了对大量字符串历史记录的支持，支持设置最大记录条数，并进行了算法优化，时间复杂度 `O(1)`
- `protocol` 该包提供了一种自定义传输协议，通过结构化的数据 `8字节类型长度 + 类型字符串 + 8字节数据长度 + 数据字节数组` 进行传输，支持直接在连接 `net.Conn` 中读取和写入
