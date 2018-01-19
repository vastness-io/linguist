package detect

import (
	log "github.com/sirupsen/logrus"
	"github.com/vastness-io/linguist/pkg/action"
	"github.com/vastness-io/linguist/pkg/linguist"
	"io"
	"os"
)

var (
	logger = log.WithField("pkg", "detect")
)

type Language struct {
	Language   string  `json:"language"`
	Percentage float64 `json:"percentage"`
}

type Languages map[string]int

func DetermineLanguages(localFolder string, names action.FileNames) ([]Language, error) {

	languages := make(Languages)
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

	return mapToLanguages(languages, overallSize), os.RemoveAll(localFolder)
}

func getFileContents(bufferSize int32, filePath string) ([]byte, error) {

	buf := make([]byte, bufferSize)

	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	_, err = f.Read(buf)

	if err = f.Close(); err != nil {
		return nil, err
	}

	if err != io.EOF {
		return nil, err
	}

	return buf, nil
}

func traverseDirectoryLogging(path string, file action.RepoFile, skipped bool, language string, args ...interface{}) {
	logger.WithFields(log.Fields{
		"path":     path,
		"file":     file,
		"skipped":  skipped,
		"language": language,
	}).Debug(args)
}

func mapToLanguages(langMap Languages, dirSize int) []Language {
	languages := make([]Language, 0)

	for lang, size := range langMap {
		languages = append(languages, Language{
			Language:   lang,
			Percentage: (float64(size) / float64(dirSize)) * 100,
		})
	}

	return languages
}
