package main

import (
	"fmt"
	"time"

	"github.com/colmpat/words-that-matter/pkg/analysis"
	"github.com/colmpat/words-that-matter/pkg/db"
)

type IngestorBuilder struct {
	watch    string
	done     string
	interval time.Duration
	db       *db.DB
}

func NewIngestorBuilder() *IngestorBuilder {
	return &IngestorBuilder{
		watch:    "watch",
		done:     "done",
		interval: 5 * time.Second,
	}
}

func (ib *IngestorBuilder) Watch(watch string) *IngestorBuilder {
	ib.watch = watch
	return ib
}

func (ib *IngestorBuilder) Done(done string) *IngestorBuilder {
	ib.done = done
	return ib
}

func (ib *IngestorBuilder) Interval(interval time.Duration) *IngestorBuilder {
	ib.interval = interval
	return ib
}

func (ib *IngestorBuilder) DB(db *db.DB) *IngestorBuilder {
	ib.db = db
	return ib
}

func (ib *IngestorBuilder) Build() (*Ingestor, error) {
	if ib.db == nil {
		return nil, fmt.Errorf("DB is required")
	}

	return &Ingestor{
		watch:        ib.watch,
		done:         ib.done,
		interval:     ib.interval,
		inProgress:   SyncSet{set: make(map[string]struct{})},
		db:           ib.db,
		strategies:   StrategyMap{m: make(map[string]analysis.AnalysisStrategy)},
		crumbChan:    make(chan Crumb),
		analysisChan: make(chan AnalysisResult),
	}, nil
}
