package main

import (
	"net/http"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"fmt"
	"strconv"
)

func apiContainers(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		SendHttpError(500, fmt.Sprintf("Docker: %v", err), w)
		return
	}

	options := types.ContainerListOptions{}
	if all, err := strconv.ParseBool(QueryGetOrDefaultValue(r, "all", "false")); err != nil {
		SendHttpError(400, fmt.Sprintf("Invalid all parameter type %v", err), w)
		return
	} else {
		options.All = all
	}

	containers, err := docker.ContainerList(context.Background(), options)
	if err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
		return
	}

	SendJSONOrError(w, containers)
}

func apiContainersDelete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		SendHttpError(500, fmt.Sprintf("Docker: %v", err), w)
		return
	}

	options := types.ContainerRemoveOptions{}
	if force, err := strconv.ParseBool(QueryGetOrDefaultValue(r, "force", "false")); err != nil {
		SendHttpError(400, fmt.Sprintf("Invalid force parameter type %v", err), w)
		return
	} else {
		options.Force = force
	}

	containerID, ok := params["id"]
	if !ok {
		SendHttpError(400, "Missing container ID", w)
		return
	}

	if err := docker.ContainerRemove(context.Background(), containerID, options); err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
		return
	}

	SendOk(w)
}
