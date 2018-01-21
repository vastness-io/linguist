package detect

import (
	log "github.com/sirupsen/logrus"
	definition "github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist/pkg/linguist"
)

var (
	logger = log.WithField("pkg", "detect")
)

type LanguageWithSize map[string]int

type FileNames []string

func DetermineLanguages(names FileNames) []*definition.Language {

	languages := make(LanguageWithSize)
	overallSize := 0

	for _, name := range names {

		path := name
		sizeInBytes := 1 // Only care about the occurrence of the file.

		traverseDirectoryLogging(
			path,
			false,
			"Working on file")

		if linguist.ShouldIgnoreFilename(path) {
			traverseDirectoryLogging(
				path,
				true,
				"",
				"Should ignore filename")

			continue

		} else {

			byFileName := linguist.LanguageByFilename(path)

			if byFileName != "" {
				traverseDirectoryLogging(
					path,
					false,
					byFileName,
					"Detected language by filename")
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

func traverseDirectoryLogging(path string, skipped bool, language string, args ...interface{}) {
	logger.WithFields(log.Fields{
		"path":     path,
		"skipped":  skipped,
		"language": language,
	}).Debug(args)
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
