package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"net/http"
	"strconv"
)

func apiVolumes(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		SendHttpError(500, fmt.Sprintf("Docker: %v", err), w)
		return
	}

	f, err := filters.FromParam(QueryGetOrDefaultValue(r, "filters", ""))
	if err != nil {
		SendHttpError(400, fmt.Sprintf("Invalid all parameter type %v", err), w)
		return
	}

	body, err := docker.VolumeList(context.Background(), f)
	if err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
		return
	}

	SendJSONOrError(w, body)
}

func apiVolumesDelete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		SendHttpError(500, fmt.Sprintf("Docker: %v", err), w)
		return
	}

	force, err := strconv.ParseBool(QueryGetOrDefaultValue(r, "force", "false"))
	if err != nil {
		SendHttpError(400, fmt.Sprintf("Invalid force parameter type %v", err), w)
		return
	}

	volumeID, ok := params["id"]
	if !ok {
		SendHttpError(400, "Missing container ID", w)
		return
	}

	if err := docker.VolumeRemove(context.Background(), volumeID, force); err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
		return
	}

	SendOk(w)
}
