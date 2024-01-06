package manabox_test

import (
	"testing"

	"github.com/senspooky/manabox-deck-scraper/pkg/manabox"
	"github.com/stretchr/testify/assert"
)

func TestNewManabox(t *testing.T) {

}

func TestGetDecklist(t *testing.T) {
	m := manabox.New(manabox.Config{
		BaseURL: "https://cloud.manabox.app",
	})
	var tests = []struct {
		name     string
		input    string
		wantDeck bool
		wantErr  error
	}{
		{
			name:     "Valid Deck",
			input:    "BfaniIYVS3Gw2S_Kuj9csw",
			wantDeck: true,
			wantErr:  nil,
		},
		{
			name:     "Deck does not Exist",
			input:    "doesnotexist",
			wantDeck: false,
			wantErr: manabox.ErrNon200Repsonse{
				StatusCode: 404,
				Body:       "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			deck, err := m.GetDecklist(tt.input)
			assert.Equal(t, tt.wantErr, err) // Check that the error is correct
			switch tt.wantDeck {
			case true:
				assert.NotNil(t, deck) // Check that the deck is not nil
			case false:
				assert.Nil(t, deck) // Check that the deck has a nil value
			}
		})
	}

}
