package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/colmpat/words-that-matter/pkg/analysis"
)

type AnalysisResult struct {
	Path       string
	Language   string
	WordCounts map[string]int
	Err        error
}

// analyze is a goroutine that analyzes files in the crumb channel and sends the results to the analysis channel. Assumes that the channel tables have already been initialized.
func (ing *Ingestor) analyze(crumb Crumb) {
	strat, err := ing.GetStrategy(crumb.Language)
	if err != nil {
		log.Fatalf("Error getting analysis strategy: %s\n", err)
	}

	res := ing.analyzeFile(strat, crumb)
	ing.cleanup(crumb.Path)
	ing.analysisChan <- res
}

func (ing *Ingestor) analyzeFile(strat analysis.AnalysisStrategy, crumb Crumb) AnalysisResult {
	res := AnalysisResult{
		Path:       crumb.Path,
		Language:   crumb.Language,
		WordCounts: make(map[string]int),
		Err:        nil,
	}

	// open the file
	file, err := os.Open(crumb.Path)
	if err != nil {
		log.Printf("Error opening file: %s\n", err)
		res.Err = err
		return res
	}
	defer file.Close()

	br := bufio.NewScanner(file)
	for br.Scan() {
		line := br.Text()
		words, err := strat.Analyze(line)
		if err != nil {
			log.Printf("Error analyzing line: %s\n", err)
			continue
		}

		for _, word := range words {
			res.WordCounts[word]++
		}
	}

	return res
}

func (ing *Ingestor) cleanup(path string) {
	// move the file to the done directory
	if _, rest, found := strings.Cut(path, ing.watch); found {
		newPath := filepath.Join(ing.done, rest)
		if err := os.MkdirAll(filepath.Dir(newPath), os.ModePerm); err != nil {
			log.Printf("Error creating directory: %s\n", err)
			return
		}

		if err := os.Rename(path, newPath); err != nil {
			log.Printf("Error moving file: %s\n", err)
			return
		}
	}

	// mark the file as done
	ing.inProgress.Lock()
	delete(ing.inProgress.set, path)
	ing.inProgress.Unlock()
}
