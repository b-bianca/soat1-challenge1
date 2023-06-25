package client

import (
	"soat1-challenge1/internal/core/ports"

	"soat1-challenge1/internal/core/domain"

	restful "github.com/emicklei/go-restful/v3"
)

type ClientHandler struct {
	clientUseCase ports.ClientUseCase
}

func NewClientHandler(clientUseCase ports.ClientUseCase, ws *restful.WebService) *ClientHandler {
	handler := &ClientHandler{
		clientUseCase: clientUseCase,
	}

	ws.Route(ws.GET("/client/{id}").To(handler.Get).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/client").To(handler.List).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.POST("/client").To(handler.Create).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))

	return handler
}

func (tdh *ClientHandler) Get(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")

	result, err := tdh.clientUseCase.Get(id)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var client *Client = &Client{}

	client.FromDomain(result)
	resp.WriteAsJson(client)
}

func (tdh *ClientHandler) Create(req *restful.Request, resp *restful.Response) {
	var data = new(domain.Client)
	if err := req.ReadEntity(data); err != nil {
		resp.WriteError(500, err)
		return
	}

	result, err := tdh.clientUseCase.Create(data.Name, data.CPF, data.Email)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var client Client = Client{}
	client.FromDomain(result)
	resp.WriteAsJson(client)
}

func (tdh *ClientHandler) List(req *restful.Request, resp *restful.Response) {
	result, err := tdh.clientUseCase.List()
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var clients ClientList = ClientList{}

	clients = clients.FromDomain(result)
	resp.WriteAsJson(clients)
}
