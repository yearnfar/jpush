package jpush

// PushResult 推送结果
type PushResult struct {
	SendNo string `json:"sendno"`
	MsgID  string `json:"msg_id"`
}

// ErrorResult 错误返回
type ErrorResult struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
