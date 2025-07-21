package debug

import (
	"fmt"

	intconfig "github.com/evsor/kubetlkt/internal/config"
	intk8s "github.com/evsor/kubetlkt/internal/k8s"
	"github.com/spf13/cobra"
)

func NewDebugCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug",
		Short: "Debug Kubernetes deployments",
	}
	cmd.AddCommand(startCmd)
	cmd.AddCommand(cleanupCmd)
	return cmd
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a Kubernetes deployment",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := intconfig.Load()
		if err != nil {
			fmt.Println("Config not found. Using default configuration.")
			cfg, err = intconfig.LoadDefault()
		}
		fmt.Println("Creating deployment...")
		err = intk8s.CreateDeployment(cfg.Repository, cfg.Image, "default")
		if err != nil {
			fmt.Println("Failed to create deployment:", err)
			return
		}
		fmt.Println("Deployment created.")
	},
}

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Delete a Kubernetes deployment",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := intconfig.Load()
		if err != nil {
			fmt.Println("Config not found. Using default configuration.")
			cfg, err = intconfig.LoadDefault()
		}
		fmt.Println("Deleting deployment...")
		err = intk8s.DeleteDeployment(cfg.Image, "default")
		if err != nil {
			fmt.Println("Failed to delete deployment:", err)
			return
		}
		fmt.Println("Deployment deleted in namespace 'default'.")
	},
}
