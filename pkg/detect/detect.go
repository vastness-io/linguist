package detect

import (
	log "github.com/sirupsen/logrus"
	"github.com/vastness-io/linguist/pkg/action"
	"github.com/vastness-io/linguist/pkg/linguist"
)

var (
	logger = log.WithField("pkg", "detect")
)

func DetermineLanguages(names action.RepoFiles) []Language {

	languages := make(LanguageWithSize)
	overallSize := 0

	for _, name := range names {

		path := name.Name
		sizeInBytes := name.Size

		traverseDirectoryLogging(
			path,
			name,
			false,
			"Working on file")

		if linguist.ShouldIgnoreFilename(path) {
			traverseDirectoryLogging(
				path,
				name,
				true,
				"",
				"Should ignore filename")

			continue

		} else {

			byFileName := linguist.LanguageByFilename(path)

			if byFileName != "" {
				traverseDirectoryLogging(
					path,
					name,
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

func traverseDirectoryLogging(path string, file action.RepoFile, skipped bool, language string, args ...interface{}) {
	logger.WithFields(log.Fields{
		"path":     path,
		"file":     file,
		"skipped":  skipped,
		"language": language,
	}).Debug(args)
}

func mapToLanguages(langMap LanguageWithSize, dirSize int) []Language {
	languages := make([]Language, 0)

	for lang, size := range langMap {
		languages = append(languages, Language{
			Language:   lang,
			Percentage: (float64(size) / float64(dirSize)) * 100,
		})
	}

	return languages
}
