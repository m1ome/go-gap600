package g600

import "encoding/json"

type TransactionConfirmationResponse struct {
	Hash          string  `json:"Hash"`
	OutputAddress string  `json:"outputAddress"`
	Username      string  `json:"username"`
	Status        string  `json:"status"`
	ScoreTime     string  `json:"scoreTime"`
	AgentID       string  `json:"agentID"`
	Size          int     `json:"size"`
	TxValueBTC    float64 `json:"txValueBTC"`
	TxValueUSD    float64 `json:"txValueUSD"`
	TxValue       float64 `json:"txValue$"`
}

type Response struct {
	Status  int             `json:"status"`
	Type    string          `json:"type,omitempty"`
	Hash    string          `json:"hash,omitempty"`
	Message json.RawMessage `json:"message,omitempty"`

	FeeInSatoshis   int64  `json:"recommendedFeeInSatoshi"`
	MinimumFeeRatio string `json:"recommendedMinimumFeeRatio"`
}
