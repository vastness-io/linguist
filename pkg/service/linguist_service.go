package service

import (
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist/pkg/language"
)

type linguistService struct {}

func NewLinguistService() LinguistService {
	return &linguistService{}
}

func (s *linguistService) CalculateLanguages(fileNames language.FileNames) []*linguist.Language {
	return language.CalculateLanguages(fileNames)
}