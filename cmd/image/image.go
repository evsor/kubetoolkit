package image

import (
	"fmt"
	intconfig "github.com/evsor/kubetlkt/internal/config"
	intdocker "github.com/evsor/kubetlkt/internal/docker"
	"github.com/spf13/cobra"
)

func NewImageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Build and push Docker images",
	}
	cmd.AddCommand(buildAndPushCmd)
	return cmd
}

var buildAndPushCmd = &cobra.Command{
	Use:   "build-and-push",
	Short: "Build and push the debug Docker image",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := intconfig.Load()
		if err != nil {
			fmt.Println("Config not found. Using default configuration.")
			cfg, err = intconfig.LoadDefault()
		}
		image := cfg.Repository + "/" + cfg.Image + ":" + cfg.Tag
		contextDir := "."
		fmt.Println("Building and pushing image:", image)
		if err := intdocker.BuildAndPush(image, contextDir); err != nil {
			fmt.Println("Docker build or push failed:", err)
			return
		}
		fmt.Println("Build and push successful.")
	},
}
