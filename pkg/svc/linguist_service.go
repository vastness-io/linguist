package svc

import (
	"context"
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist/pkg/detect"
)

type LinguistService struct{}

func (s *LinguistService) GetLanguages(ctx context.Context, req *linguist.LanguageRequest) (*linguist.LanguageResponse, error) {

	return &linguist.LanguageResponse{
		Language: detect.DetermineLanguages(req.FileNames),
	}, nil

}
