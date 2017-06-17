package main

import (
	"net/http"
	"log"
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func apiContainers(w http.ResponseWriter, r *http.Request, params map[string]string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		log.Panicf("Docker client: %v", err)
	}

	containers, err := docker.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Panicf("Docker client: %v", err)
	}

	for _, container := range containers {
		log.Printf("Container: %v", strings.Join(container.Names, ","))
	}
}

func apiContainersDelete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to containers remove")
}
