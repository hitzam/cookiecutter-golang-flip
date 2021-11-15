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

	goCoreLog "gitlab.com/flip-id/go-core/helpers/log"
	goCoreHttp "gitlab.com/flip-id/go-core/http"
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
	ApiKey  string
}

const (
	USER_INFO_ENDPOINT = "/v1/user/info"
)

//Client clevertap client struct
type Client struct {
	BaseURL    string
	ApiKey     string
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
		ApiKey:     option.ApiKey,
		HTTPClient: goCoreHttp.NewHttpClient(nil),
	}
}

//SendEventData will send event
func (ct *Client) GetUserInfo(ctx context.Context, jwtToken string) (err error) {
	logger := goCoreLog.GetLogger(ctx)

	req, err := ct.newRequest(http.MethodGet, USER_INFO_ENDPOINT, jwtToken. nil)
	if err != nil {
		logger.FormatLog("create new request", err, jwtToken).Error("failed")
		return err
	}

	_, err = ct.do(req)
	if err != nil {
		return err
	}

	return nil
}

func (ct *Client) newRequest(method, path string, jwtToken string, body interface{}) (*http.Request, error) {
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
	req.Header.Set("api-key", ct.ApiKey)
	req.Header.Set("Authorization", "Bearer "+jwtToken)

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
