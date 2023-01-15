package backend

import (
	"strconv"

	"github.com/emicklei/go-restful/v3"
)

func WithListOptionsParam(builder *restful.RouteBuilder) *restful.RouteBuilder {
	return builder.Param(restful.QueryParameter("page_size", "The maximum number of return items")).
		Param(restful.QueryParameter("page_num", "The first (page_size * page_num) items will be skipped.")).
		Param(restful.QueryParameter("filter", "The returned items will be filtered according to the filter."))
}

func ListOptionsFromRequest(req *restful.Request) (opts ListOptions, err error) {
	if raw := req.QueryParameter("page_size"); raw != "" {
		opts.Limit, err = strconv.ParseInt(req.QueryParameter("page_size"), 10, 64)
		if err != nil {
			return
		}
	}
	if raw := req.QueryParameter("page_num"); raw != "" {
		opts.Offset, err = strconv.ParseInt(req.QueryParameter("page_num"), 10, 64)
		if err != nil {
			return
		}
	}
	opts.Filter = req.QueryParameter("filter")
	return
}

type ListOptions struct {
	Limit  int64
	Offset int64
	Filter string
}
