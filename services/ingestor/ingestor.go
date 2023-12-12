package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/colmpat/words-that-matter/pkg/analysis"
	"github.com/uptrace/bun"
)

type Ingestor struct {
	watch      string
	done       string
	interval   time.Duration
	inProgress SyncSet
	db         *bun.DB

	strategies StrategyMap

	crumbChan    chan Crumb
	analysisChan chan AnalysisResult
}

type StrategyMap struct {
	sync.Mutex
	m map[string]analysis.AnalysisStrategy
}

type SyncSet struct {
	sync.Mutex
	set map[string]struct{}
}

func (ing *Ingestor) Start() {
	// create the watch and done directories if they don't exist
	if err := os.MkdirAll(ing.watch, 0755); err != nil {
		log.Fatalf("Error creating watch directory: %s\n", err)
	}
	if err := os.MkdirAll(ing.done, 0755); err != nil {
		log.Fatalf("Error creating done directory: %s\n", err)
	}

	go func() {
		// if ingest() returns, it means there was an error with the watcher
		// so we start it again
		for {
			ing.ingest()
		}
	}()

	go ing.upload()
}

// GetStrategy returns the strategy for the given language.
func (ing *Ingestor) GetStrategy(language string) (analysis.AnalysisStrategy, error) {
	ing.strategies.Lock()
	defer ing.strategies.Unlock()

	// check if we have already built the strategy for this language
	strat, ok := ing.strategies.m[language]
	if ok {
		return strat, nil
	}

	strat, err := analysis.GetStrategy(language)
	if err == nil {
		ing.strategies.m[language] = strat
	}

	return strat, err
}
