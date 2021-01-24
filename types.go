package payze_go

import (
	"fmt"
	"net/http"
)

type Callbacks struct {
	CallbackSuccess string `json:"callbackError,omitempty"`
	CallbackFail    string `json:"callback,omitempty"`
}

/**PayZeRequest PayZe general request structure*/
type PayZeRequest struct {
	Method    string      `json:"method"`
	APIKey    string      `json:"apiKey"`
	APISecret string      `json:"apiSecret"`
	Data      interface{} `json:"data"`
}

/**PayRequestData Data part of PayZe general requests*/
type PayRequestData struct {
	Amount       float64 `json:"amount,string"`
	Currency     string  `json:"currency"`
	Preauthorize bool    `json:"preauthorize,omitempty"`
	Callbacks
}

/**PayWithCardRequestData Data part of PayZe general requests for card payment*/
type PayWithCardRequestData struct {
	Amount       float64 `json:"amount,string"`
	Currency     string  `json:"currency"`
	Preauthorize bool    `json:"preauthorize,omitempty"`
	CardToken    string  `json:"card_token"`
}

type SplitParticipant struct {
	IBan   string  `json:"iban"`
	Amount float64 `json:"amount,string"`
	PayIn  int     `json:"payIn"`
}

/**ResponseField response part of PayZe general requests*/
type ResponseField struct {
	CardID         string `json:"cardId,omitempty"`
	TransactionUrl string `json:"transactionUrl,omitempty"`
	TransactionID  string `json:"transactionId,omitempty"`
}

/**PayZeResponse PayZe general response structure*/
type PayZeResponse struct {
	ID        int64         `json:"id"`
	Status    string        `json:"status"`
	Action    string        `json:"action"`
	CreatedAt string        `json:"createdDate"`
	Response  ResponseField `json:"response"`
}

type ErrorResponse struct {
	Response  *http.Response `json:"-"`
	Data      []byte         `json:"-"`
	Timestamp string         `json:"timestamp"`
	Status    uint           `json:"status"`
	ErrorHint string         `json:"error"`
	Message   string         `json:"message"`
	Path      string         `json:"path"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Timestamp: %v\nStatus: %v\nError: %v\nMessage: %v\nPath: %v\n",
		e.Timestamp, e.Status, e.ErrorHint, e.Message, e.Path)
}
