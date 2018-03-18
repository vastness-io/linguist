package http

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

// ServiceLevelFunc is responsible for handling the decoded request and returning a "response" and possibly an error.
type ServiceLevelFunc func(ctx context.Context, req interface{}) (res interface{}, err error)

// DecodeRequestFunc is responsible for decoding the request which will be passed to the ServiceLevelFunc.
type DecodeRequestFunc func(context.Context, *http.Request) (req interface{}, err error)

// EncodeResponseFunc is responsible for encoding the response from the serviceLevelFunc and returning to caller.
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{})

// ErrorEncoderFunc is responsible for handing the error returned in any step in the flow.
type ErrorEncoderFunc func(context.Context, http.ResponseWriter, error)

type handler struct {
	svcLevelFunc   ServiceLevelFunc
	decodeReqFunc  DecodeRequestFunc
	encodeResFunc  EncodeResponseFunc
	errEncoderFunc ErrorEncoderFunc
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	decodedReq, err := h.decodeReqFunc(ctx, req)

	if err != nil {
		h.errEncoderFunc(ctx, w, err)
		return
	}

	res, err := h.svcLevelFunc(ctx, decodedReq)

	if err != nil {
		h.errEncoderFunc(ctx, w, err)
		return
	}

	h.encodeResFunc(ctx, w, res)
}

// NewHandler creates a http.Handler which executes a encoder/decoder for request/response, error handling and serviceLevel function.
func NewHandler(decoder DecodeRequestFunc, encoder EncodeResponseFunc, errorHandlerFunc ErrorEncoderFunc, serviceLevelFunc ServiceLevelFunc) http.Handler {
	return &handler{
		svcLevelFunc:   serviceLevelFunc,
		decodeReqFunc:  decoder,
		encodeResFunc:  encoder,
		errEncoderFunc: errorHandlerFunc,
	}

}

// HTTPRoute is a helper struct for http handlers.
type HTTPRoute struct {
	Path    string
	Methods []string
	Handler http.Handler
}

// NewHTTPRouter creates a Router with the registered http handlers.
func NewHTTPRouter(httpRoutes ...HTTPRoute) *mux.Router {
	r := mux.NewRouter()

	for _, route := range httpRoutes {
		r.Handle(route.Path, route.Handler).Methods(route.Methods...)
	}

	return r
}
