package main

import (
	"fmt"

	dockercli "github.com/docker/docker/client"
)

func main() {
	_, _ = dockercli.NewClientWithOpts()
	fmt.Println("docker is crazy")
}
