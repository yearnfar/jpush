package jpush

// Message 自定义消息
// 应用内消息。或者称作：自定义消息，透传消息。
// 此部分内容不会展示到通知栏上，JPush SDK 收到消息内容后透传给 App。需要 App 自行处理。
// iOS 平台上，此部分内容在推送应用内消息通道（非APNS）获取。Windows Phone 暂时不支持应用内消息。
type Message struct {
	MsgContent string                 `json:"msg_content"`       // 消息内容本身
	Title      string                 `json:"title,omitempty"`   // 消息标题
	Content    string                 `json:"content,omitempty"` // 消息内容类型
	Extras     map[string]interface{} `json:"extras,omitempty"`  // JSON 格式的可选参数
}
