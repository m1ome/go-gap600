package g600

import (
	"errors"
	"net/http"
)

var (
	// DefaultURL URL to be used in every request
	DefaultURL = "https://account.gap600.com"

	// DefaultClient is HTTP client to be user in every client
	DefaultClient = &http.Client{}

	// DefaultAgent API Agent in gap600
	DefaultAgent = "Primary"

	// StatusOk 200 is a "good" status in gap600
	StatusOk = 200
)

var (
	// ErrEmptyApiKey thrown when you not specified api key
	ErrEmptyApiKey = errors.New("no api key provided")

	// ErrTransactionNotFound thrown when there is no such transaction
	ErrTransactionNotFound = errors.New("transaction not found")

	// ErrOutputNotFound throws when there is no such input
	ErrOutputNotFound = errors.New("output not found")
)
