package analysis

import (
	"regexp"
	"strings"

	"github.com/go-ego/gse"
)

// EnglishStrategy is an analysis strategy for English
type EnglishStrategy struct {
	segmenter gse.Segmenter
}

// NewEnglishStrategy returns a new EnglishStrategy
func NewEnglishStrategy() (*EnglishStrategy, error) {
	es := &EnglishStrategy{}
	es.segmenter.LoadDict()
	return es, nil
}

// Analyze returns a slice of stemmed words from the given line
func (es *EnglishStrategy) Analyze(line string) ([]string, error) {
	segments := es.segmenter.Segment([]byte(line))
	words := make([]string, 0, len(segments))
	for _, seg := range segments {
		clean := Clean(seg.Token().Text())
		clean = strings.TrimSpace(clean)
		if len(clean) > 0 {
			words = append(words, clean)
		}
	}

	return words, nil
}

// Clean removes all non-alphanumeric characters from the line
func Clean(line string) string {
	// return everything that isn't a unicode letter or number
	return regexp.MustCompile(`[^\p{L}\p{N}]+`).ReplaceAllString(line, " ")
}
