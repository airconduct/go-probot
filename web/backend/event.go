package backend

import (
	"context"
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func AddEventService(ws *restful.WebService, es EventSource) {
	handler := &eventHandler{es: es}
	tags := []string{"events"}
	ws.Route(WithListOptionsParam(
		ws.GET("/events").To(handler.ListEvent).
			Metadata(restfulspec.KeyOpenAPITags, tags).
			Doc("get all events").
			Returns(200, "OK", EventList{}),
	))
}

type eventHandler struct {
	es EventSource
}

func (h *eventHandler) ListEvent(req *restful.Request, resp *restful.Response) {
	ctx, cancel := context.WithCancel(req.Request.Context())
	defer cancel()
	opts, err := ListOptionsFromRequest(req)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}

	events, err := h.es.ListEvent(ctx, opts)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	if err := resp.WriteEntity(events); err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
}

type EventSource interface {
	ListEvent(ctx context.Context, opts ListOptions) (EventList, error)
}

type EventList struct {
	Items []Event
}

type Event struct {
	// Name of event, it should be shown on dashboard
	Name string `json:"name"`
	// Type of event. All events can be grouped by the type.
	Type EventType `json:"type"`
	// Metrics is the metrical data about this event
	Metrics EventMetrics `json:"metrics"`
}

type EventType string

type EventMetrics struct {
	Count int64 `json:"count"`
}
