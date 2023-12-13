package db

import (
	"github.com/colmpat/words-that-matter/pkg/models"
	"github.com/uptrace/bun"
)

// GetDeckWithTotals returns a Deck with the totals for the given word counts
func GetDeckWithTotals(wordCount map[string]int) *models.Deck {
	totalTerms := 0
	uniqueTerms := len(wordCount)
	uniqueTermsUsedOnce := 0

	for _, count := range wordCount {
		totalTerms += count
		if count == 1 {
			uniqueTermsUsedOnce++
		}
	}

	return &models.Deck{
		TotalTerms:          totalTerms,
		UniqueTerms:         uniqueTerms,
		UniqueTermsUsedOnce: uniqueTermsUsedOnce,
	}
}

// UploadMedia uploads a media file to the database with the given metadata and word counts. Does so in a transaction for atomicity.
func (db *DB) UploadMedia(metadata models.Media, wordCount map[string]int) error {
	return db.Transact(func(tx *bun.Tx) error {
		// create deck
		deck := GetDeckWithTotals(wordCount)
		_, err := tx.NewInsert().
			Model(deck).
			Exec(db.Ctx)
		if err != nil {
			return err
		}

		terms := make([]models.Term, 0, len(wordCount))
		frequencies := make([]models.Frequency, 0, len(wordCount))

		// create frequencies
		for word, count := range wordCount {
			terms = append(terms, models.Term{
				Term:     word,
				Language: metadata.Language,
				Count:    count,
			})
			frequencies = append(frequencies, models.Frequency{
				Count:    count,
				Term:     word,
				Language: metadata.Language,
				DeckID:   deck.ID,
			})
		}

		// update/insert terms
		_, err = tx.NewInsert().
			Model(&terms).
			On("CONFLICT (term, language) DO UPDATE").
			Set("count = Term.count + EXCLUDED.count").
			Exec(db.Ctx)
		if err != nil {
			return err
		}

		// then insert frequencies
		_, err = tx.NewInsert().
			Model(&frequencies).
			Exec(db.Ctx)
		if err != nil {
			return err
		}

		// then insert media
		metadata.DeckID = deck.ID
		_, err = tx.NewInsert().
			Model(&metadata).
			Exec(db.Ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
