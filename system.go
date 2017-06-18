package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"net/http"
)

func apiSystemInfo(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		SendHttpError(500, fmt.Sprintf("Docker: %v", err), w)
		return
	}

	body, err := docker.Info(context.Background())
	if err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
		return
	}

	SendJSONOrError(w, body)
}
