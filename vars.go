package g600

import (
	"net/http"
	"errors"
)

var (
	// DefaultURL URL to be used in every request
	DefaultURL    = "https://account.gap600.com"

	// DefaultClient is HTTP client to be user in every client
	DefaultClient = &http.Client{}

	// DefaultAgent API Agent in gap600
	DefaultAgent  = "Primary"

	// StatusOK 200 is a "good" status in gap600
	StatusOk      = 200
)

var (
	// ErrEmptyApiKey thrown when you not specified api key
	ErrEmptyApiKey  = errors.New("no api key provided")
)
