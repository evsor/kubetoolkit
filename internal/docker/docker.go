package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	dockercfg "github.com/cpuguy83/dockercfg"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
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

func Push(imageName string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	ctx := context.Background()

	username, password, err := dockercfg.GetRegistryCredentials("https://index.docker.io/v1/")
	if err != nil {
		return fmt.Errorf("failed to load docker config: %w", err)
	}
	authConfig := registry.AuthConfig{
		Username: username,
		Password: password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	resp, err := cli.ImagePush(ctx, imageName, image.PushOptions{RegistryAuth: authStr})
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
