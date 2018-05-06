package language

import (
	definition "github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist/third_party/linguist"
	"strings"
)

type Language map[string]*definition.Language

type FileNames []string

func CalculateLanguages(names FileNames) []*definition.Language {

	var (
		languages           = make(Language)
		overallSize float64 = 0
	)

	for _, name := range names {

		var (
			path                 = name
			sizeInBytes  float64 = 1 // Only care about the occurrence of the file.
			santizedPath         = removeDirectoryPrefix(path)
		)

		if linguist.ShouldIgnoreFilename(santizedPath) {
			continue
		}
		byFileName := linguist.LanguageByFilename(santizedPath)

		if byFileName != "" {
			_, exists := languages[byFileName]
			if !exists {
				languages[byFileName] = &definition.Language{
					FileNames: make([]string, 0),
				}
			}

			languages[byFileName].Percentage += sizeInBytes
			languages[byFileName].FileNames = append(languages[byFileName].GetFileNames(), path)
			overallSize += sizeInBytes
			continue
		}

		_, exists := languages["Other"]

		if !exists {
			languages["Other"] = &definition.Language{
				FileNames: make([]string, 0),
			}
		}

		languages["Other"].Percentage += float64(sizeInBytes)
		languages["Other"].FileNames = append(languages["Other"].GetFileNames(), path)
		overallSize += sizeInBytes
	}

	return mapToLanguages(languages, overallSize)
}

func mapToLanguages(langMap Language, dirSize float64) []*definition.Language {
	languages := make([]*definition.Language, 0)

	for langName, lang := range langMap {
		languages = append(languages, &definition.Language{
			Name:       langName,
			Percentage: (lang.GetPercentage() / dirSize) * 100,
			FileNames:  lang.GetFileNames(),
		})
	}

	return languages
}

func removeDirectoryPrefix(file string) string {

	index := strings.LastIndex(file, "/")

	if index != -1 {
		r := []rune(file)
		return string(r[index+1:])
	}

	return file
}
