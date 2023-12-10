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

		// todo - upload to db
		fmt.Printf("Uploaded %s\n", meta.Name)
	}
}
