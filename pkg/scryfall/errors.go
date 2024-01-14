package scryfall

import (
	"fmt"
	"strings"
)

type Error struct {
	Object   string   `json:"object"`
	Code     string   `json:"code"`
	Status   int      `json:"status"`
	Warnings []string `json:"warnings"`
	Details  string   `json:"details"`
}

func (e Error) Error() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("error making request to scryfall (HTTP %d): %s", e.Status, e.Details))
	for _, w := range e.Warnings {
		sb.WriteString(fmt.Sprintf("\n\t%s", w))
	}
	return sb.String()
}
