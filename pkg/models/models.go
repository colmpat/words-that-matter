package models

import (
	"github.com/uptrace/bun"
)

type Account struct {
	// should be unique over provider and provider_account_id
	bun.BaseModel `bun:"table:Account,unique:provider,provider_account_id"`

	ID                int     `bun:"id,pk,autoincrement" json:"id"`
	Type              string  `json:"type"`
	Provider          string  `json:"provider"`
	AccessToken       string  `json:"access_token"`
	ProviderAccountId string  `json:"provider_account_id"`
	RefreshToken      *string `json:"refresh_token"`
	ExpiresAt         *int    `json:"expires_at"`
	TokenType         *string `json:"token_type"`
	Scope             *string `json:"scope"`
	IdToken           *string `json:"id_token"`
	SessionState      *string `json:"session_state"`
	UserID            int     `json:"user_id"`

	User *User `bun:"rel:has-one,join:user_id=id" json:"user"`
}

type User struct {
	bun.BaseModel `bun:"table:User"`

	ID      int    `bun:"id,pk,autoincrement" json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Image   string `json:"image"`
	IsAdmin bool   `bun:",default:false"`
}

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
