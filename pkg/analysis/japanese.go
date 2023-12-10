package analysis

import (
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// JapaneseStrategy is an AnalysisStrategy that uses the kagome tokenizer to tokenize Japanese text
type JapaneseStrategy struct {
	t *tokenizer.Tokenizer
}

// NewJapaneseStrategy returns a new JapaneseStrategy
func NewJapaneseStrategy() (*JapaneseStrategy, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}
	return &JapaneseStrategy{t}, nil
}

// Analyze returns a slice of words from the given line
func (js *JapaneseStrategy) Analyze(line string) ([]string, error) {
	return js.t.Wakati(line), nil
}
