package server

import (
	"context"
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist/pkg/service"
)

type LinguistServer struct {
	LinguistService service.LinguistService
}

func NewLinguistServer(linguistService service.LinguistService) *LinguistServer {
	return &LinguistServer{
		LinguistService: linguistService,
	}
}

func (s *LinguistServer) GetLanguages(ctx context.Context, req *linguist.LanguageRequest) (*linguist.LanguageResponse, error) {
	return &linguist.LanguageResponse{
		Language: s.LinguistService.CalculateLanguages(req.FileNames),
	}, ctx.Err()

}
