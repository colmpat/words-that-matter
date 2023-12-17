package db

import (
	"github.com/colmpat/words-that-matter/pkg/models"
)

func (db *DB) GetMedia() ([]models.Media, error) {
	models := []models.Media{}
	err := db.NewSelect().
		Model(&models).
		Scan(db.Ctx)

	return models, err
}
