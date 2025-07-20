package main

import (
	"github.com/evsor/kubetlkt/cmd/config"
	"github.com/evsor/kubetlkt/cmd/debug"
	"github.com/evsor/kubetlkt/cmd/image"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kubetlkt",
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
