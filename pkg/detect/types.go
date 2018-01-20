package detect

type Language struct {
	Language   string  `json:"language"`
	Percentage float64 `json:"percentage"`
}

type LanguageWithSize map[string]int
