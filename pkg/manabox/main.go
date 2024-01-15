package manabox

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type Manabox interface {
	GetDecklist(string) (*Deck, error)
}

type Config struct {
	BaseURL string
}

type manabox struct {
	h *resty.Client
	c *Config
}

func New(c Config) Manabox {
	m := &manabox{}
	m.c = &c
	m.h = resty.New().
		SetBaseURL(c.BaseURL).
		SetTimeout(10 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		OnError(func(req *resty.Request, err error) {
			if v, ok := err.(*resty.ResponseError); ok {
				v.Response.Request.Error = ErrNon200Repsonse{
					StatusCode: v.Response.StatusCode(),
					Body:       string(v.Response.Body()),
				}
			}
			// Log the error, increment a metric, etc...
		})
	return m
}

func (m *manabox) GetDecklist(id string) (*Deck, error) {
	deck := &Deck{}
	_, err := m.h.R().
		SetBody(deck).
		SetPathParam("id", id).
		Get("/decks/{id}")
	if err != nil {
		// nil better than zero; zero is technically valid.
		return nil, err
	}
	return deck, err
}
