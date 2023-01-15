package backend

import (
	restful "github.com/emicklei/go-restful/v3"
)

func WebService(event EventSource) *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON) // you can specify this per route as well

	AddEventService(ws, event)
	return ws
}
