package detect

import (
	log "github.com/sirupsen/logrus"
	"github.com/vastness-io/linguist/pkg/linguist"
	"io"
	"os"
	"path/filepath"
)

var (
	logger = log.WithField("pkg", "detect")
)

type Language struct {
	Language   string  `json:"language"`
	Percentage float64 `json:"percentage"`
}

type Languages map[string]int64

func DetermineLanguages(parentPath string, bufferSize int32) ([]Language, error) {

	languages := make(Languages)
	var overallSize int64 = 0

	filepath.Walk(parentPath, func(path string, file os.FileInfo, err error) error {

		fileSize := file.Size()

		traverseDirectoryLogging(
			path,
			file,
			false,
			"Working on file")

		if linguist.ShouldIgnoreFilename(path) {
			traverseDirectoryLogging(
				path,
				file,
				true,
				"",
				"Should ignore filename")

			if file.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if file.IsDir() {

			if file.Name() == ".git" {
				traverseDirectoryLogging(
					path,
					file,
					true,
					"Skipping .git folder")
				return filepath.SkipDir
			}

		} else {
			if linguist.ShouldIgnoreFilename(path) {
				traverseDirectoryLogging(
					path,
					file,
					true,
					"",
					"Should ignore filename")
				return nil
			}

			byFileName := linguist.LanguageByFilename(path)

			if byFileName != "" {
				traverseDirectoryLogging(
					path,
					file,
					false,
					byFileName,
					"Detected language by filename")
				languages[byFileName] += fileSize
				overallSize += fileSize
				return nil
			}

			content, err := getFileContents(bufferSize, path)

			if err != nil {
				logger.Error(err)
				return filepath.SkipDir
			}

			if linguist.ShouldIgnoreContents(content) {
				traverseDirectoryLogging(
					path,
					file,
					true,
					"",
					"Should ignore contents")
				return filepath.SkipDir
			}

			hints := linguist.LanguageHints(path)

			determinedLang := linguist.LanguageByContents(content, hints)

			traverseDirectoryLogging(
				path,
				file,
				false,
				"hints", hints,
				"Detected language by file contents")

			if determinedLang != "" {
				languages[determinedLang] += fileSize
				overallSize += fileSize
				return nil
			}

			languages["Other"] += fileSize
			overallSize += fileSize

		}
		return nil
	})

	return mapToLanguages(languages, overallSize), os.RemoveAll(parentPath)
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

func traverseDirectoryLogging(path string, file os.FileInfo, skipped bool, language string, args ...interface{}) {
	logger.WithFields(log.Fields{
		"path":     path,
		"file":     file,
		"skipped":  skipped,
		"language": language,
	}).Debug(args)
}

func mapToLanguages(langMap Languages, dirSize int64) []Language {
	languages := make([]Language, 0)

	for lang, size := range langMap {
		languages = append(languages, Language{
			Language:   lang,
			Percentage: (float64(size) / float64(dirSize)) * 100,
		})
	}

	return languages
}
