package db

import (
	"github.com/google/uuid"
)

// Represents a deck in the database
type Deck struct {
	// gorm.Model

	Name   string
	Format Format

	Cards  []Card
	Source ExternalRef // The deck that this deck is synced from
	// TODO: How will decks be synced?
	// SyncTo   []ExternalRef // The decks that this deck is synced to
}

// Represents an external reference to a deck
type ExternalRef struct {
	// gorm.Model

	// ! DeckID and Source are expected to be unique for a given Deck
	DeckID string
	Source Source
}

// Represents a card in a deck the database
type Card struct {
	// gorm.Model

	// ! DeckID, BoardCategory, and CardData's ID are expected to be unique for a given Card in a deck
	DeckID        uint
	BoardCategory BoardCategory
	CardData      CardData

	Name string
}

// Represents a card
type CardData struct {
	// gorm.Model

	// ! ID is expected to be unique for a given Card in a deck
	ID   uuid.UUID // Scryfall ID
	Name string
	// TODO: Add more fields
}
