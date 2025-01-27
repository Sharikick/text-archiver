package handler

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "mycli",
    Short: "My CLI application",
    Long: "This is my first CLI app with cobra.",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Error")
		os.Exit(1)
	}
}
