package cli

import (
	"sync"

	"github.com/spf13/cobra"
)

var (
	once    sync.Once
	Root    *cobra.Command
	Serve   *cobra.Command
	Connect *cobra.Command
	Auth    *string
	Host    *string
	Port    *string
)
