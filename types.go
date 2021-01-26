package payze_go

import (
	"fmt"
	"net/http"
)

type Callbacks struct {
	CallbackSuccess string `json:"callback,omitempty"`
	CallbackFail string `json:"callbackError,omitempty"`
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

/**SplitInfo is an entry of splits array returned during action: GetTransactionInfo*/
type SplitInfo struct {
	IBan   string  `json:"iban"`
	Amount float64 `json:"amount,string"`
	Status string 	`json:"status,omitempty"`
	//CashOutOrder string `json:"cashOutOrder,omitempty"`
}

/**LogEntry is an entry of logs array returned during action: GetTransactionInfo*/
type LogEntry struct {
	Date string `json:"date"`//TODO for future parse dates
	StatusBefore string `json:"statusBefore"`
	Status string `json:"status"`
	ChangeType string `json:"changeType"`
}

/**ResponseField response part of PayZe general requests*/
type ResponseField struct {
	CardID         string `json:"cardId,omitempty"`
	TransactionUrl string `json:"transactionUrl,omitempty"`
	TransactionID  string `json:"transactionId,omitempty"`
}

/**ResponseField response part of PayZe general requests*/
type TransactionResponseField struct {
	CardMask         string `json:"cardMask,omitempty"`
	Amount			 float64 `json:"amount,omitempty"`
	Splits			 []SplitInfo `json:"splits,omitempty"`
	Log  			 []LogEntry `json:"log,omitempty"`
	Err 			 string `json:"error,omitempty"`
	GetCanBeCommitted bool `json:"getCanBeCommitted"`
	ResultCode 		  string `json:"resultCode"`
	TransactionID	string `json:"transactionId"`
	CommitDate 		string `json:"commit_date"`
	FinalAmount		float64 `json:"finalAmount"`
	Currency		string `json:"currency"`
	Commission		float64 `json:"commission"`
	Refundable		bool `json:"refundable"`
	Refunded 		float64 `json:"refunded"`
	HasSplit		bool `json:"hasSplit"`
	Status			string `json:"status"`
	CreatedAt		string `json:"createDate"`
}

/**PayZeResponse PayZe general response structure*/
type PayZeResponse struct {
	ID        int64         `json:"id"`
	Status    string        `json:"status"`
	Action    string        `json:"action"`
	CreatedAt string        `json:"createdDate"`
	Response  ResponseField `json:"response"`
}

/**PayZeTransactionInfoResponse PayZe response structure of transaction info*/
type PayZeTransactionInfoResponse struct {
	ID        int64         `json:"id"`
	Status    string        `json:"status"`
	Action    string        `json:"action"`
	CreatedAt string        `json:"createdDate"`
	Response  TransactionResponseField `json:"response"`
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
