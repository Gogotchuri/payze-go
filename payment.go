package payze_go

func (c *Client) JustPay(paymentData PayRequestData) (*PayZeResponse, error) {
	req, err := c.NewRequest(DefaultHTTPMethod, c.APIBase, c.CreateRequestPayload(JustPay, paymentData))
	if err != nil {
		return nil, err
	}
	resp := &PayZeResponse{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CreateSubscription(subData SubscriptionRequestData) (*PayZeResponse, error) {
	req, err := c.NewRequest(DefaultHTTPMethod, c.APIBase, c.CreateRequestPayload(CreateSubscription, subData))
	if err != nil {
		return nil, err
	}
	resp := &PayZeResponse{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) JustPayAndSplit(amount float64, commitNow bool, currency string, split []SplitParticipant) error {
	panic("Not implemented")
}

func (c *Client) AddCard(callbacks Callbacks) (*PayZeResponse, error) {
	req, err := c.NewRequest(DefaultHTTPMethod, c.APIBase, c.CreateRequestPayload(AddCard, callbacks))
	if err != nil {
		return nil, err
	}
	resp := &PayZeResponse{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) PayWithCard(paymentData PayWithCardRequestData) (*PayZeResponse, error) {
	req, err := c.NewRequest(DefaultHTTPMethod, c.APIBase, c.CreateRequestPayload(PayWithCard, paymentData))
	if err != nil {
		return nil, err
	}
	resp := &PayZeResponse{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) PayWithCardAndSplit() {
	panic("Not implemented")

}

func (c *Client) GetTransactionInfo(transactionID string) (*PayZeTransactionInfoResponse, error) {
	transIDT := struct {
		TransactionID string `json:"transactionId"`
	}{TransactionID: transactionID}

	req, err := c.NewRequest(DefaultHTTPMethod, c.APIBase, c.CreateRequestPayload(TransactionInfo, transIDT))
	if err != nil {
		return nil, err
	}
	resp := &PayZeTransactionInfoResponse{}
	err = c.Send(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) RefundTransaction() {
	panic("Not implemented")

}

func (c *Client) GetMerchantBalance() {
	panic("Not implemented")

}

func (c *Client) CommitTransaction() {
	panic("Not implemented")

}
