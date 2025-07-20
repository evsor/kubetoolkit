package image

import (
	"fmt"
	intconfig "github.com/evsor/kubetoolkit/internal/config"
	intdocker "github.com/evsor/kubetoolkit/internal/docker"
	"github.com/spf13/cobra"
)

func NewImageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Build and push Docker images",
	}
	cmd.AddCommand(buildCmd)
	cmd.AddCommand(pushCmd)
	return cmd
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the Docker image",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := intconfig.Load()
		if err != nil {
			fmt.Println("Config not found. Please run 'kubetoolkit config init' first.")
			return
		}
		image := cfg.Repository + "/" + cfg.Image + ":latest"
		contextDir := "."
		fmt.Println("Building image:", image)
		if err := intdocker.Build(image, contextDir); err != nil {
			fmt.Println("Docker build failed:", err)
			return
		}
		fmt.Println("Build successful.")
	},
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push the Docker image",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := intconfig.Load()
		if err != nil {
			fmt.Println("Config not found. Please run 'kubetoolkit config init' first.")
			return
		}
		image := cfg.Repository + "/" + cfg.Image + ":latest"
		fmt.Println("Pushing image:", image)
		if err := intdocker.Push(image); err != nil {
			fmt.Println("Docker push failed:", err)
			return
		}
		fmt.Println("Push successful.")
	},
}
