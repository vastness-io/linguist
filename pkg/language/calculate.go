package language

import (
	definition "github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist/third_party/linguist"
)

type LanguageWithSize map[string]int

type FileNames []string

func CalculateLanguages(names FileNames) []*definition.Language {

	languages := make(LanguageWithSize)
	overallSize := 0

	for _, name := range names {

		path := name
		sizeInBytes := 1 // Only care about the occurrence of the file.

		if linguist.ShouldIgnoreFilename(path) {
			continue
		} else {

			byFileName := linguist.LanguageByFilename(path)

			if byFileName != "" {
				languages[byFileName] += sizeInBytes
				overallSize += sizeInBytes
				continue
			}

			languages["Other"] += sizeInBytes
			overallSize += sizeInBytes
		}
	}

	return mapToLanguages(languages, overallSize)
}

func mapToLanguages(langMap LanguageWithSize, dirSize int) []*definition.Language {
	languages := make([]*definition.Language, 0)

	for lang, size := range langMap {
		languages = append(languages, &definition.Language{
			Name:       lang,
			Percentage: (float64(size) / float64(dirSize)) * 100,
		})
	}

	return languages
}
