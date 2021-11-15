package flipserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	goCoreHttp "gitlab.com/flip-id/go-core/http"
	goCoreLog "gitlab.com/flip-id/go-core/helpers/log"
)

//SendMessageCallbackRequest user profile payload
type SendMessageCallbackRequest struct {
	MessageId   string     `json:"message_id"`
	Status      string     `json:"status"`
	PhoneNumber string     `json:"phone_number"`
	Timestamp   *time.Time `json:"timestamp"`
}

//Option clevertap option
type Option struct {
	BaseUrl string
}

const (
	CALLBACK_ENDPOINT = "/v2/...." //Not implemented yet
)

//Client clevertap client struct
type Client struct {
	BaseURL    string
	HTTPClient *goCoreHttp.HttpClient
}

//IFlipServerClient interface
type IFlipServerClient interface {
	SendMessageCallBack(ctx context.Context, payload SendMessageCallbackRequest) (err error)
}

//NewCleverTapClient instantiate
func NewFlipServerClient(option Option) IFlipServerClient {
	return &Client{
		BaseURL:    option.BaseUrl,
		HTTPClient: goCoreHttp.NewHttpClient(nil),
	}
}

//SendEventData will send event
func (ct *Client) SendMessageCallBack(ctx context.Context, payload SendMessageCallbackRequest) (err error) {
	logger := goCoreLog.GetLogger(ctx)

	req, err := ct.newRequest(http.MethodPost, CALLBACK_ENDPOINT, payload)
	if err != nil {
		logger.FormatLog("create new request", err, CALLBACK_ENDPOINT).Error("failed")
		return err
	}

	_, err = ct.do(req)
	if err != nil {
		return err
	}

	return nil
}

func (ct *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	u := fmt.Sprintf("%s%s", ct.BaseURL, path)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return req, nil
}

func (ct *Client) do(req *http.Request) (bodyByte []byte, err error) {
	resp, err := ct.HTTPClient.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyByte, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		if resp.StatusCode != 200 {
			err := fmt.Errorf("client response: %s", string(bodyByte))
			return nil, err
		}
	}

	return bodyByte, nil
}
