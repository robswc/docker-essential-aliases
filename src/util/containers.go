package util

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Container struct {
	Id       string
	Name     string
	Accessor string
}

// returns a container
func dockerInspect(idOrName string) Container {
	container := Container{}
	dockerCmd := exec.Command("docker", "inspect", idOrName)
	output, err := dockerCmd.Output()
	if err != nil {
		if strings.Contains(string(output), "[]") {
			return container
		}
	} else {
		var containers []Container
		if err := json.Unmarshal(output, &containers); err != nil {
			panic(err)
		}
		container = containers[0]
	}
	container.Accessor = container.Id
	return container
}

func CaptureContainer(containerNameOrId string) {

	container := dockerInspect(containerNameOrId)

	// save the container id, name to environment variables
	SaveToTemp(WorkingContainerId, container.Id)
	SaveToTemp(WorkingContainerName, container.Name)

}

func GetContainer(providedNameOrId string) Container {

	container := Container{}

	if providedNameOrId != "" {
		container = dockerInspect(providedNameOrId)
		if container.Id == "" {
			fmt.Println("Container: " + providedNameOrId + " not found.")
			os.Exit(1)
		}
	} else {
		fmt.Println("No container name or id provided. Using the last container.")
		// see if we can use the last container ID first
		container = dockerInspect(GetFromTemp(WorkingContainerId))
		if container.Id == "" {
			// if not, try the last container name
			container = dockerInspect(GetFromTemp(WorkingContainerName))
			if container.Id == "" {
				fmt.Println("No container found. Please provide a container name or id.")
				os.Exit(1)
			}
		}
	}

	return container
}
