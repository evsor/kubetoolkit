package docker

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func Build(image, contextDir string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	ctx := context.Background()

	tar, err := archive.TarWithOptions(contextDir, &archive.TarOptions{})
	if err != nil {
		return err
	}
	defer tar.Close()

	buildOptions := types.ImageBuildOptions{
		Tags:       []string{image},
		Remove:     true,
		Dockerfile: "Dockerfile",
	}

	resp, err := cli.ImageBuild(ctx, tar, buildOptions)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	return nil
}

func Push(image string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	ctx := context.Background()
	resp, err := cli.ImagePush(ctx, image, image.PushOptions{})
	if err != nil {
		return err
	}
	defer resp.Close()
	io.Copy(os.Stdout, resp)
	return nil
}

func BuildAndPush(image, contextDir string) error {
	if err := Build(image, contextDir); err != nil {
		return fmt.Errorf("docker build failed: %w", err)
	}
	if err := Push(image); err != nil {
		return fmt.Errorf("docker push failed: %w", err)
	}
	return nil
}
