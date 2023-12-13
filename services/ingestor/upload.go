package main

import (
	"fmt"

	"github.com/colmpat/words-that-matter/pkg/models"
)

func (ing *Ingestor) upload() {
	for res := range ing.analysisChan {
		fmt.Printf("Uploading %s...\n", res.Path)

		meta := models.Media{
			Name:     res.Path,
			Language: res.Language,
			Type:     string(models.MediaTypeMovie),
		}

		err := ing.db.UploadMedia(meta, res.WordCounts)
		if err != nil {
			fmt.Printf("Error uploading %s: %s\n", res.Path, err)
		}
	}
}
