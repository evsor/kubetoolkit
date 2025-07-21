package config

import (
	"fmt"
	"strings"

	"github.com/evsor/kubetlkt/internal/config"
	"github.com/spf13/cobra"
)

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage kubetlkt configuration",
	}
	cmd.AddCommand(setImageCmd)
	return cmd
}

var setImageCmd = &cobra.Command{
	Use:   "set-image repository/image tag",
	Short: "Set the debug Docker image (format: repository/image tag)",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		repoImage := args[0]
		tag := args[1]
		parts := strings.SplitN(repoImage, "/", 2)
		if len(parts) != 2 || tag == "" {
			fmt.Println("Usage: set-image repository/image tag")
			return
		}
		repo := parts[0]
		image := parts[1]
		cfg := config.Config{
			Repository: repo,
			Image:      image,
			Tag:        tag,
		}
		err := config.Save(&cfg)
		if err != nil {
			fmt.Println("Failed to save config:", err)
			return
		}
		fmt.Println("Config updated at:")
		path, _ := config.ConfigFilePath()
		fmt.Println(path)
	},
}
