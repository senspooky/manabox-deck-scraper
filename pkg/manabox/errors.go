package manabox

import (
	"fmt"
)

// complex errors
type ErrNon200Repsonse struct {
	StatusCode int
	Body       string
}

func (e ErrNon200Repsonse) Error() string {
	return fmt.Sprintf("bad status recieved: %d\n%s", e.StatusCode, e.Body)
}
