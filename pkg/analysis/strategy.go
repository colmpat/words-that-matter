package analysis

import (
	"fmt"
)

// the AnalysisStrategy interface is used to analyze a line of text and return a slice of words for a specific language
type AnalysisStrategy interface {
	Analyze(line string) ([]string, error)
}

// GetStrategy returns an AnalysisStrategy for the given language
func GetStrategy(language string) (AnalysisStrategy, error) {
	switch language {
	case "english":
		return NewEnglishStrategy()
	case "japanese":
		return NewJapaneseStrategy()
	default:
		return nil, fmt.Errorf("unknown language: %s", language)
	}
}

// ValidLanguage returns true if the given language is supported
func ValidLanguage(language string) bool {
	switch language {
	case "english", "japanese":
		return true
	default:
		return false
	}
}
