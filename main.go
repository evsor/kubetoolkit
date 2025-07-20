package main

import (
	"github.com/evsor/kubetoolkit/cmd/config"
	"github.com/evsor/kubetoolkit/cmd/debug"
	"github.com/evsor/kubetoolkit/cmd/image"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kubetoolkit",
	Short: "Kubernetes workloads debugging tool.",
}

func main() {
	rootCmd.AddCommand(config.NewConfigCommand())
	rootCmd.AddCommand(image.NewImageCommand())
	rootCmd.AddCommand(debug.NewDebugCommand())
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
