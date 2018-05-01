package jpush

import "net/http"

const (
	API = "https://%s@%s:api.jpush.cn/v3/push"
)

type Client struct {
	AppKey    string
	AppSecret string
}

func send() {

	c := http.Client{}
	c.BaseAuth()
}
