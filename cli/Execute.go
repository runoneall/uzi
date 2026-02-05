package cli

import (
	"log/slog"
)

func Execute() {
	if err := Root.Execute(); err != nil {
		slog.Error(err.Error())
	}
}
