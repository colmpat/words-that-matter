package models

import (
	"github.com/uptrace/bun"
)

// Term model
type Term struct {
	bun.BaseModel `bun:"table:Term"`

	Term     string `bun:",pk" json:"term"`
	Language string `bun:",pk" json:"language"`
	Count    int    `json:"count"`

	Frequencies []Frequency `bun:"rel:has-many" json:"frequencies"`
}

// Frequency model
type Frequency struct {
	bun.BaseModel `bun:"table:Frequency"`

	ID       int    `bun:"id,pk,autoincrement" json:"id"`
	Count    int    `json:"count"`
	Term     string `json:"term_id"`
	Language string `json:"language"`
	DeckID   int    `json:"deck_id"`

	TermModel *Term `bun:"rel:has-one,join:term=term,join:language=language" json:"term_model"`
	DeckModel *Deck `bun:"rel:has-one,join:deck_id=id" json:"deck_model"`
}

// Deck model
type Deck struct {
	bun.BaseModel `bun:"table:Deck"`

	ID                  int `bun:"id,pk,autoincrement" json:"id"`
	TotalTerms          int `bun:",default:0" json:"total_terms"`
	UniqueTerms         int `bun:",default:0" json:"unique_terms"`
	UniqueTermsUsedOnce int `bun:",default:0" json:"unique_terms_used_once"`

	Frequencies []Frequency `bun:"rel:has-many" json:"frequencies"`
}

// Media model
type Media struct {
	bun.BaseModel `bun:"table:Media"`

	ID          int    `bun:"id,pk,autoincrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Year        int    `json:"year"`
	IMDbID      string `bun:"imdb_id" json:"imdb_id"`
	DeckID      int    `json:"deck_id"`
	Language    string `json:"language"`

	DeckModel *Deck `bun:"rel:has-one,join:deck_id=id" json:"deck_model"`
}

// MediaType enum
type MediaType string

const (
	MediaTypeTV    MediaType = "TV"
	MediaTypeMovie MediaType = "MOVIE"
	MediaTypeBook  MediaType = "BOOK"
)
