package language

import (
	"github.com/vastness-io/linguist-svc"
	"reflect"
	"testing"
)

func TestDetermineLanguages(t *testing.T) {
	var (
		repoFiles = FileNames{

			"main.go",
			"pom.xml",
			"build.xml",
			"README.md",
			"invalid_file_format",
		}
	)
	calculateLanguagesTests := []struct {
		fileNames         FileNames
		expectedLanguages []*linguist.Language
	}{
		{
			fileNames: repoFiles,
			expectedLanguages: []*linguist.Language{
				{
					Name:       "Go",
					Percentage: float64(25),
					FileNames: []string{
						"main.go",
					},
				},
				{
					Name:       "Maven POM",
					Percentage: float64(25),
					FileNames: []string{
						"pom.xml",
					},
				},
				{
					Name:       "Ant Build System",
					Percentage: float64(25),
					FileNames: []string{
						"build.xml",
					},
				},
				{
					Name:       "Other",
					Percentage: float64(25),
					FileNames: []string{
						"invalid_file_format",
					},
				},
			},
		},
	}

	for _, test := range calculateLanguagesTests {

		langs := CalculateLanguages(test.fileNames)

		if !reflect.DeepEqual(test.expectedLanguages, langs) {
			t.Errorf("expected: %v, got: %v", test.expectedLanguages, langs)
		}
	}
}

func TestRemoveDirectoryPrefix(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "i have no parent",
			out: "i have no parent",
		},
		{
			in:  "parent/pom.xml",
			out: "pom.xml",
		},
		{
			in:  "1/2/3/4/more/depth/pom.xml",
			out: "pom.xml",
		},
		{
			in:  "",
			out: "",
		},
		{
			in:  "x/",
			out: "",
		},
		{
			in:  "//",
			out: "",
		},
		{
			in:  "pom.xml",
			out: "pom.xml",
		},
	}

	for _, test := range tests {

		result := removeDirectoryPrefix(test.in)
		if !reflect.DeepEqual(test.out, result) {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}

}
