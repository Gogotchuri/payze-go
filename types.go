package payze_go

import (
	"fmt"
	"net/http"
)

type CallbackURLs struct {
	CallbackSuccess string `json:"callbackError"`
	CallbackFail    string `json:"callback"`
}

type ErrorResponse struct {
	Response	*http.Response	`json:"-"`
	Data 		[]byte			`json:"-"`
	//TODO don't know error response structure for now, but its fields should be added and response should be parsed into them
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%v", e.Data)
}

