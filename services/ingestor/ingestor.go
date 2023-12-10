package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/uptrace/bun"
)

type Ingestor struct {
	watch      string
	done       string
	interval   time.Duration
	inProgress RWSet
	db         *bun.DB

	// channel table for communicating with the analysis workers (one channel per language)
	crumbChanTable map[string]chan Crumb
	analysisChan   chan AnalysisResult
}

type RWSet struct {
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

// get the channel for a given language and start the analysis worker if it's not already running
func (ing *Ingestor) GetCrumbChan(lang string) chan Crumb {
	c, ok := ing.crumbChanTable[lang]
	if !ok {
		c = make(chan Crumb)
		ing.crumbChanTable[lang] = c

		go ing.analyze(lang)
	}

	return c
}

type IngestorBuilder struct {
	watch    string
	done     string
	interval time.Duration
	db       *bun.DB
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

func (ib *IngestorBuilder) DB(db *bun.DB) *IngestorBuilder {
	ib.db = db
	return ib
}

func (ib *IngestorBuilder) Build() (*Ingestor, error) {
	if ib.db == nil {
		return nil, fmt.Errorf("DB is required")
	}

	return &Ingestor{
		watch:          ib.watch,
		done:           ib.done,
		interval:       ib.interval,
		inProgress:     RWSet{set: make(map[string]struct{})},
		db:             ib.db,
		crumbChanTable: make(map[string]chan Crumb),
		analysisChan:   make(chan AnalysisResult),
	}, nil
}
