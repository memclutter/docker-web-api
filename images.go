package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"net/http"
	"strconv"
)

func apiImages(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		SendHttpError(500, fmt.Sprintf("Docker: %v", err), w)
		return
	}

	options := types.ImageListOptions{}
	if all, err := strconv.ParseBool(QueryGetOrDefaultValue(r, "all", "false")); err != nil {
		SendHttpError(400, fmt.Sprintf("Invalid all parameter type %v", err), w)
		return
	} else {
		options.All = all
	}

	if f, err := filters.FromParam(QueryGetOrDefaultValue(r, "filters", "")); err != nil {
		SendHttpError(400, fmt.Sprintf("Invalid filters %v", err), w)
		return
	} else {
		options.Filters = f
	}

	images, err := docker.ImageList(context.Background(), options)
	if err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
		return
	}

	SendJSONOrError(w, images)
}

func apiImagesDelete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		SendHttpError(500, fmt.Sprintf("Docker: %v", err), w)
		return
	}

	options := types.ImageRemoveOptions{}
	if force, err := strconv.ParseBool(QueryGetOrDefaultValue(r, "force", "false")); err != nil {
		SendHttpError(400, fmt.Sprintf("Invalid force parameter type %v", err), w)
		return
	} else {
		options.Force = force
	}

	imageID, ok := params["id"]
	if !ok {
		SendHttpError(400, "Missing image ID", w)
		return
	}

	if body, err := docker.ImageRemove(context.Background(), imageID, options); err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
	} else {
		SendJSONOrError(w, body)
	}
}
