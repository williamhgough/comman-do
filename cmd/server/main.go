package main

import (
	"net/http"

	"github.com/williamhgough/comman-do/internal/commandoserver"
	"github.com/williamhgough/comman-do/rpc/todo"
)

func main() {
	server := &commandoserver.Server{}
	twirpHandler := todo.NewCommandoServer(server, nil)

	http.ListenAndServe(":8080", twirpHandler)
}
