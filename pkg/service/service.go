package service

import (
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist/pkg/language"
)

type LinguistService interface {
	CalculateLanguages(language.FileNames) []*linguist.Language
}
