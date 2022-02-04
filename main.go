package main

import (
	dockercli "github.com/docker/docker/client"
)

func main() {
	_, _ = dockercli.NewClientWithOpts()
	print("docker is crazy")
}
