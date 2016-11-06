package worker

import (
	"context"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

type DockerSandbox struct {
	client      *client.Client
	image       string
	logs        string
	containerId string
}

// BUG(#2) We should not try to download the image every time we start a new container. We should check if it exists locally.
func NewDockerSandbox(image string) (*DockerSandbox, error) {
	s := new(DockerSandbox)
	err := (error)(nil)
	s.image = image
	s.logs = ""

	s.client, err = client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	err = s.downloadImage()
	if err != nil {
		return nil, err
	}

	err = s.createContainer()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *DockerSandbox) Destroy() {
}

func (s *DockerSandbox) Run(command string, connection bool) {
}

func (s *DockerSandbox) GetLogs() string {
	return s.logs
}

// We will create a container
// This part is a bit hacky because we will run a
// command that hangs but do nothing.
// It will prevent the container from stopping
// and will enable us to run our desired commands...
func (s *DockerSandbox) createContainer() error {
	container, err := s.client.ContainerCreate(
		context.Background(),
		&(container.Config{
			Cmd:   []string{"/bin/sh", "-c", "while true; do sleep 86400; done"},
			Image: s.image,
		}),
		&(container.HostConfig{}),
		&(network.NetworkingConfig{}),
		"") // We don't want a name for our container

	s.containerId = container.ID
	return err
}

//BUG(#1) Logs are not checked for error when pulling an image
func (s *DockerSandbox) downloadImage() error {
	reader, err := s.client.ImageCreate(
		context.Background(),
		s.image,
		types.ImageCreateOptions{},
	)

	if err != nil {
		s.logs += "Unable to pull the image " + s.image + "\n"
		return err
	}

	bytesRead, err := ioutil.ReadAll(reader)
	defer reader.Close()
	if err != nil {
		s.logs += "Unable to read logs when pulling the image " + s.image + "\n"
		return err
	}
	s.logs += string(bytesRead[:])

	return nil
}
