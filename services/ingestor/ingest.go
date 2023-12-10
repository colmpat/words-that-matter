package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/colmpat/words-that-matter/pkg/analysis"
)

// Crumb is a file that has been found in the watch directory... ready to be ingested.
type Crumb struct {
	Path     string
	Language string
}

func (ing *Ingestor) ingest() {
	ticker := time.NewTicker(ing.interval)
	defer ticker.Stop()

	for range ticker.C {
		ing.inProgress.Lock()
		crumbs, err := ing.walkWatchFiles()
		ing.inProgress.Unlock()
		if err != nil {
			fmt.Printf("Error walking watch directory: %s\n", err)
			continue
		}

		for _, crumb := range crumbs {
			ing.GetCrumbChan(crumb.Language) <- crumb
		}
	}
}

// Walk the watch directory, returns a list of files to process. If a file is not valid, it is moved straight to the done directory. Updates the inProgress set.
func (ing *Ingestor) walkWatchFiles() ([]Crumb, error) {
	var crumbs []Crumb

	err := filepath.WalkDir(ing.watch, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			_, inProg := ing.inProgress.set[path]

			dir := filepath.Dir(path)
			lang := filepath.Base(dir)
			valid := analysis.ValidLanguage(lang)

			fileExt := strings.ToLower(filepath.Ext(path))
			validExt := fileExt == ".txt" || fileExt == ".text"

			if !inProg && valid && validExt {
				crumbs = append(crumbs, Crumb{Path: path, Language: lang})
				ing.inProgress.set[path] = struct{}{}
			}
		}

		return nil
	})

	return crumbs, err
}
