package scryfall

import (
	"fmt"
	"time"

	"github.com/imroc/req/v3"
)

type Config struct {
	BaseURL     string
	LastRequest time.Time // This is used to slow requests as specified by docs
}

type Scryfall interface {
}

type scryfall struct {
	h *req.Client
	c *Config
}

func New(c Config) Scryfall {
	s := &scryfall{}
	s.c = &c
	s.h = req.
		SetBaseURL(c.BaseURL).
		//EnableDumpAll().
		SetCommonErrorResult(&Error{}).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil { // There is an underlying error
				return nil
			}
			if err, ok := resp.ErrorResult().(*Error); ok {
				resp.Err = err // Convert api error into go error
				return nil
			}
			if !resp.IsSuccessState() {
				resp.Err = Error{
					Status:  resp.StatusCode,
					Details: fmt.Sprintf("raw: %s", resp.Dump()),
				}
			}
			return nil
		})
	return s
}
