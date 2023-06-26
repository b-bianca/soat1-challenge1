package main

import (
	"flag"

	"soat1-challenge1/helpers"
	"soat1-challenge1/internal/core/ports"
	"soat1-challenge1/internal/core/usecases"
	handlerClient "soat1-challenge1/internal/handlers/client"
	repoClient "soat1-challenge1/internal/repositories/client"

	restful "github.com/emicklei/go-restful/v3"
)

var (
	repo    string
	binding string
)

func init() {
	flag.StringVar(&repo, "repo", "mysql", "Mongo or MySql")
	flag.StringVar(&binding, "httpbind", ":8080", "address/port to bind listen socket")

	flag.Parse()
}

func main() {
	var clientRepo ports.ClientRepository

	clientRepo = startMysqlRepo()

	clientUseCase := usecases.NewClientUseCase(clientRepo)

	ws := new(restful.WebService)
	ws = ws.Path("/api")
	handlerClient.NewClientHandler(clientUseCase, ws)
	restful.Add(ws)

	//log.Info("Listening...")

	//log.Panic(http.ListenAndServe(binding, nil))
}

func startMysqlRepo() ports.ClientRepository {
	return repoClient.NewClientMysqlRepo(helpers.StartMysqlDb())
}
