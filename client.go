package jpush

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Client 客户端
type Client struct {
	http.Client

	Username string
	Password string
}

// NewClient 实例化 Client
func NewClient(username string, password string) *Client {
	return &Client{Username: username, Password: password}
}

// 解析返回数据
func (c *Client) parseResponse(response *http.Response, result interface{}) error {
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		errResult := &ErrorResult{}
		err = json.Unmarshal(data, errResult)
		if err != nil {
			log.Printf("解析数据失败: %v", err)
			return errors.New("接口返回失败")
		}

		return fmt.Errorf("接口错误：%s", errResult.Error.Message)
	}

	return json.Unmarshal(data, result)
}

// PostJSON Post请求
func (c *Client) PostJSON(urlStr string, param interface{}, result interface{}) (err error) {
	body, err := json.Marshal(param)
	if err != nil {
		return
	}

	fmt.Println(string(body))

	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewReader(body))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	fmt.Println(req.Header.Get("Authorization"))

	resp, err := c.Do(req)
	if err != nil {
		return
	}

	return c.parseResponse(resp, result)
}
