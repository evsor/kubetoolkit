package cmd

import (
	"fmt"
	"os"

	"github.com/evsor/kubetlkt/cmd/config"
	"github.com/evsor/kubetlkt/cmd/debug"
	"github.com/evsor/kubetlkt/cmd/image"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(debug.DebugCmd)
	rootCmd.AddCommand(image.ImageCmd)
}

var rootCmd = &cobra.Command{
	Use:   "kubetlkt",
	Short: "Kubernetes workloads debugging tool.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
