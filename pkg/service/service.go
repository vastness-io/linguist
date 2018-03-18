package service

import (
	"github.com/vastness-io/linguist/pkg/language"
	"github.com/vastness-io/linguist-svc"
)

type LinguistService interface {
	CalculateLanguages(language.FileNames) []*linguist.Language
}