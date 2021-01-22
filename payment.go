package payze_go

//
//
//type SplitParticipant struct {
//	IBan   string  `json:"iban"`
//	Amount float64 `json:"amount,string"`
//	PayIn  int     `json:"payIn"`
//}
//
//type payZeRequestT struct {
//	Method    string      `json:"method"`
//	APIKey    string      `json:"apiKey"`
//	APISecret string      `json:"apiSecret"`
//	Data      interface{} `json:"data"`
//}
//
//func (c *Client) JustPay(amount float64, commitNow bool, currency string) error {
//	data := struct {
//		Amount float64 `json:"amount,string"`
//		CommitNow bool `json:"commitNow"`
//		Currency string `json:"currency"`
//		Callback string `json:"callback"`
//		CallbackError string `json:"callbackError"`
//	}{ Amount: amount, CommitNow: commitNow, Currency: currency,
//		Callback: pz.CallbackSuccess, CallbackError: pz.CallbackFail}
//	res, err := pz.performRequest("POST", data)
//	if err != nil { return err }
//	fmt.Println(string(res.([]byte)))
//	//TODO Add struct after a request body
//	return nil
//}
//
//
//func (c *Client) JustPayAndSplit(amount float64, commitNow bool, currency string, split []SplitParticipant) error {
//	data := struct {
//		Amount float64 `json:"amount,string"` //TODO check all the types, which should be a string and which number
//		CommitNow bool `json:"commitNow"`
//		Currency string `json:"currency"`
//		Callback string `json:"callback"`
//		CallbackError string `json:"callbackError"`
//	}{ Amount: amount, CommitNow: commitNow, Currency: currency,
//		Callback: pz.CallbackSuccess, CallbackError: pz.CallbackFail}
//	res, err := pz.performRequest("POST", data)
//	if err != nil { return err }
//	fmt.Println(string(res.([]byte)))
//	//TODO Add struct after a request body
//	return nil
//}
//
//
//func (c *Client) SaveCard()  {
//
//}
//
//func (c *Client) PayWithCard()  {
//
//}
//
//func (c *Client) PayWithCardAndSplit()  {
//
//}
//
//func (c *Client) GetTransactionInfo()  {
//
//}
//
//func (c *Client) performRequest(method string, payloadData interface{}) (interface{}, error){
//	paymentPayload, err := json.Marshal(payZeRequestT{
//		Method:    method,
//		APIKey:    pz.APIKey,
//		APISecret: pz.APISecret,
//		Data:      payloadData,
//	})
//
//	client := &http.Client{}
//	req, err := http.NewRequest(method, pz.APIUrl, bytes.NewReader(paymentPayload))
//
//	if err != nil { return nil, err }
//	res, err := client.Do(req)
//	if err != nil { return nil, err }
//	defer res.Body.Close()
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil { return nil, err }
//	//parsedResp := json.Unmarshal(body)
//	return body, nil
//}