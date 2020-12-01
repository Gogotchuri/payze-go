package payze_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PayZe struct {
	APIKey          string `json:"api_key"`
	APISecret       string `json:"api_secret"`
	APIUrl          string `json:"api_url"`
	CallbackFail    string `json:"callback_fail"`
	CallbackSuccess string `json:"callback_success"`
}

type SplitParticipant struct {
	IBan   string  `json:"iban"`
	Amount float64 `json:"amount,string"`
	PayIn  int     `json:"payIn"`
}

type payZeRequestT struct {
	Method    string      `json:"method"`
	APIKey    string      `json:"apiKey"`
	APISecret string      `json:"apiSecret"`
	Data      interface{} `json:"data"`
}

func (pz *PayZe) JustPay(amount float64, commitNow bool, currency string) error {
	data := struct {
		Amount float64 `json:"amount,string"`
		CommitNow bool `json:"commitNow"`
		Currency string `json:"currency"`
		Callback string `json:"callback"`
		CallbackError string `json:"callbackError"`
	}{ Amount: amount, CommitNow: commitNow, Currency: currency,
		Callback: pz.CallbackSuccess, CallbackError: pz.CallbackFail}
	res, err := pz.performRequest("POST", data)
	if err != nil { return err }
	fmt.Println(string(res.([]byte)))
	//TODO Add struct after a request body
	return nil
}


func (pz *PayZe) JustPayAndSplit(amount float64, commitNow bool, currency string, split []SplitParticipant) error {
	data := struct {
		Amount float64 `json:"amount,string"` //TODO check all the types, which should be a string and which number
		CommitNow bool `json:"commitNow"`
		Currency string `json:"currency"`
		Callback string `json:"callback"`
		CallbackError string `json:"callbackError"`
	}{ Amount: amount, CommitNow: commitNow, Currency: currency,
		Callback: pz.CallbackSuccess, CallbackError: pz.CallbackFail}
	res, err := pz.performRequest("POST", data)
	if err != nil { return err }
	fmt.Println(string(res.([]byte)))
	//TODO Add struct after a request body
	return nil
}


func (pz *PayZe) SaveCard()  {

}

func (pz *PayZe) PayWithCard()  {

}

func (pz *PayZe) PayWithCardAndSplit()  {

}

func (pz *PayZe) GetTransactionInfo()  {

}

func (pz *PayZe) performRequest(method string, payloadData interface{}) (interface{}, error){
	paymentPayload, err := json.Marshal(payZeRequestT{
		Method:    method,
		APIKey:    pz.APIKey,
		APISecret: pz.APISecret,
		Data:      payloadData,
	})

	client := &http.Client{}
	req, err := http.NewRequest(method, pz.APIUrl, bytes.NewReader(paymentPayload))

	if err != nil { return nil, err }
	res, err := client.Do(req)
	if err != nil { return nil, err }
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil { return nil, err }
	//parsedResp := json.Unmarshal(body)
	return body, nil
}