package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"net/http"
)

func apiNetworks(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		SendHttpError(500, fmt.Sprintf("Docker: %v", err), w)
		return
	}

	options := types.NetworkListOptions{}
	if f, err := filters.FromParam(QueryGetOrDefaultValue(r, "filters", "")); err != nil {
		SendHttpError(400, fmt.Sprintf("Invalid all parameter type %v", err), w)
		return
	} else {
		options.Filters = f
	}

	body, err := docker.NetworkList(context.Background(), options)
	if err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
		return
	}

	SendJSONOrError(w, body)
}

func apiNetworksDelete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		SendHttpError(500, fmt.Sprintf("Docker: %v", err), w)
		return
	}

	networkID, ok := params["id"]
	if !ok {
		SendHttpError(400, "Missing network ID", w)
		return
	}

	if err := docker.NetworkRemove(context.Background(), networkID); err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
		return
	}

	SendOk(w)
}
