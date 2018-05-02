package jpush

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Client 客户端
type Client struct {
	http.Client

	AppKey    string
	AppSecret string
}

// NewClient 实例化 Client
func NewClient(appKey string, appSecret string) *Client {
	return &Client{AppKey: appKey, AppSecret: appSecret}
}

// 解析返回数据
func (c *Client) parseResponse(response *http.Response, result interface{}) error {
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, result)
}

// PostJSON Post请求
func (c *Client) PostJSON(urlStr string, param interface{}, result interface{}) (err error) {
	body, err := json.Marshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewReader(body))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.AppKey, c.AppSecret)

	resp, err := c.Do(req)
	if err != nil {
		return
	}

	return c.parseResponse(resp, result)
}
