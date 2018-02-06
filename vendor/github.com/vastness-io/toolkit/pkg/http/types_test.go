package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandlerExecution(t *testing.T) {

	tests := []struct {
		encode     EncodeResponseFunc
		decode     DecodeRequestFunc
		errFunc    ErrorEncoderFunc
		s          ServiceLevelFunc
		statusCode int
	}{
		{
			encode: func(_ context.Context, w http.ResponseWriter, res interface{}) {
				w.WriteHeader(200)
			},
			decode: func(context.Context, *http.Request) (req interface{}, err error) {
				return "example", nil
			},

			errFunc: func(_ context.Context, w http.ResponseWriter, err error) {
				w.WriteHeader(500)
			},

			s: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				str := req.(string)
				return str, nil
			},
			statusCode: 200,
		},
		{
			encode: func(_ context.Context, w http.ResponseWriter, res interface{}) {
				w.WriteHeader(200)
			},
			decode: func(context.Context, *http.Request) (req interface{}, err error) {
				return struct{}{}, nil
			},

			errFunc: func(_ context.Context, w http.ResponseWriter, err error) {
				// this is called
				w.WriteHeader(500)
			},

			s: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				//fail on service level func
				return nil, errors.New("some_error")
			},
			statusCode: 500,
		},
		{
			encode: func(_ context.Context, w http.ResponseWriter, res interface{}) {
				w.WriteHeader(200)
			},
			decode: func(context.Context, *http.Request) (req interface{}, err error) {
				//fail on decode
				return struct{}{}, errors.New("some_error")
			},

			errFunc: func(_ context.Context, w http.ResponseWriter, err error) {
				w.WriteHeader(500)
			},

			s: func(ctx context.Context, req interface{}) (res interface{}, err error) {

				return nil, nil
			},
			statusCode: 500,
		},
	}

	for _, test := range tests {

		r := NewHTTPRouter(HTTPRoute{
			Path:    "/",
			Methods: []string{"GET"},
			Handler: NewHandler(test.decode, test.encode, test.errFunc, test.s),
		})

		srv := httptest.NewServer(r)

		resp, _ := http.Get(srv.URL)

		srv.Close()

		if test.statusCode != resp.StatusCode {
			t.Fatalf("expected: %v, got %v", test.statusCode, resp.StatusCode)
		}

	}

}
