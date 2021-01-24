package payze_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"sync"
)

/** Client represents a PayZe REST API Client */
type Client struct {
	sync.Mutex
	Client   *http.Client
	ClientID string
	Secret   string
	APIBase  string
	Log      io.Writer // If user sets log file logs will go there
	//Default callbacks are used in case callbacks arent passed for each request data
	DefaultCallbacks Callbacks
}

/** NewClient returns new Client struct APIBase is a base API URL*/
func NewClient(clientID string, secret string, APIBase string) (*Client, error) {
	if clientID == "" || secret == "" || APIBase == "" {
		return nil, errors.New("ClientID, Secret and APIBase are required to create a Client")
	}

	return &Client{
		Client:   &http.Client{},
		ClientID: clientID,
		Secret:   secret,
		APIBase:  APIBase,
	}, nil
}

/** SetHTTPClient sets *http.Client as current client client*/
func (c *Client) SetHTTPClient(client *http.Client) {
	c.Client = client
}

/** SetLog will set/change the logging destination. */
func (c *Client) SetLog(log io.Writer) {
	c.Log = log
}

/** Send makes a request to the API, the response body will be unmarshalled into v,
which should be correct struct for the given request body passed by reference or
it can be an io.Writer, in which case the response bytes will be written to it */
func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")

	resp, err = c.Client.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		data, err = ioutil.ReadAll(resp.Body)
		errResp := &ErrorResponse{Response: resp, Data: data}
		if err == nil && len(data) > 0 {
			err = json.Unmarshal(data, errResp)
		}
		return errResp
	}
	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		io.Copy(w, resp.Body)
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

/** NewRequest constructs a new http.Request, Marshal payload to json bytes */
func (c *Client) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequest(method, url, buf)
}

/**CreateRequestPayload Create proper payload for PayZe request*/
func (c *Client) CreateRequestPayload(action string, data interface{}) PayZeRequest {
	return PayZeRequest{
		Method:    action,
		APIKey:    c.ClientID,
		APISecret: c.Secret,
		Data:      data,
	}
}

/** log will dump request and response to the log destination */
func (c *Client) log(r *http.Request, resp *http.Response) {
	if c.Log != nil {
		var (
			reqDump  string
			respDump []byte
		)

		if r != nil {
			reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
		}
		if resp != nil {
			respDump, _ = httputil.DumpResponse(resp, true)
		}

		c.Log.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	}
}
