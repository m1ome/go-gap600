package g600

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Client provides basic
type Client struct {
	opts Options
}

// New creates new Client. You should provide API key. You can find it here: https://account.gap600.com/keys.
// By default Agent ID is "Primary"
func New(opts Options) (*Client, error) {
	opts.init()

	if opts.ApiKey == "" {
		return nil, ErrEmptyApiKey
	}

	return &Client{
		opts: opts,
	}, nil
}

// TransactionConfirm checks and returns any information about transaction.
//
// If transaction don't persists in networks it will send back and error
// Two types of error can be:
// ErrTransactionNotFound - for unknown transaction
// ErrOutputNotFound - for unknown output
//
// Please, make sure you correctly reacts on this errors
func (c Client) TransactionConfirm(hash, address string) (*TransactionConfirmationResponse, error) {
	tcr := &TransactionConfirmationResponse{}

	url := fmt.Sprintf("/g600/api/v1/%s/%s/%s/%s", c.opts.ApiKey, c.opts.AgentID, hash, address)
	if _, err := c.request(url, tcr); err != nil {
		return nil, err
	}

	return tcr, nil
}

// RecommendedFee returns recommended tx fees in satoshis and fee ratio as a string
func (c Client) RecommendedFee() (int64, string, error) {
	res := &Response{}

	url := fmt.Sprintf("/g600/api/v1/%s/current-recommended-fee-ratio", c.opts.ApiKey)
	res, err := c.request(url, res)
	if err != nil {
		return 0, "", err
	}

	return res.FeeInSatoshis, res.MinimumFeeRatio, nil
}

// TransactionConfirmTestnet do same thing as TransactionConfirm but only for a testnet.
func (c Client) TransactionConfirmTestnet(hash, address string) (*TransactionConfirmationResponse, error) {
	tcr := &TransactionConfirmationResponse{}

	url := fmt.Sprintf("/g600/apitest/v1/%s/%s/%s/%s", c.opts.ApiKey, c.opts.AgentID, hash, address)
	if _, err := c.request(url, tcr); err != nil {
		return nil, err
	}

	return tcr, nil
}

func (c Client) request(url string, v interface{}) (*Response, error) {
	req, err := http.NewRequest("GET", c.opts.URL+url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.opts.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	message := &Response{}
	if err := json.NewDecoder(res.Body).Decode(message); err != nil {
		return nil, err
	}

	if message.Status != StatusOk {
		switch message.Type {
		case "TransactionNotFoundError":
			return nil, ErrTransactionNotFound
		case "OutputAddressNotFoundError":
			return nil, ErrOutputNotFound
		}

		var gapError string
		if err := json.Unmarshal(message.Message, &gapError); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("gap600 error [%d]: %s", message.Status, gapError)
	}

	if len(message.Message) > 0 {
		if err := json.Unmarshal(message.Message, v); err != nil {
			return nil, err
		}
	}

	return message, nil
}
