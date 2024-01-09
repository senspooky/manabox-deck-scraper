package manabox

import (
	"github.com/imroc/req/v3"
)

type Manabox interface {
	GetDecklist(string) (*Deck, error)
}

type Config struct {
	BaseURL string
}

type manabox struct {
	c *req.Client
}

func New(c Config) Manabox {
	m := &manabox{}
	m.c = req.
		SetBaseURL(c.BaseURL).
		//EnableDumpAll().
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil { // There is an underlying error
				return nil
			}
			if !resp.IsSuccessState() {
				resp.Err = ErrNon200Repsonse{
					StatusCode: resp.StatusCode,
					Body:       resp.String(),
				}
			}
			return nil
		})
	return m
}

func (m *manabox) GetDecklist(id string) (*Deck, error) {
	deck := &Deck{}
	_, err := m.c.R().
		SetPathParam("id", id).
		SetSuccessResult(deck).
		Get("/decks/{id}")
	if err != nil {
		// nil better than zero; zero is technically valid.
		return nil, err
	}
	return deck, err
}
