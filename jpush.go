package jpush

const (
	PushAPI = "https://api.jpush.cn/v3/push"
)

// JPushApp 极光推送
type JPushApp struct {
	AppKey    string
	AppSecret string
}

// PushAndroid 推送安卓设备
func (jp *JPushApp) PushAndroid(audience *Audience, notification *NotificationAndroid) (result *PushResult, err error) {

	return
}

// PushAndroidTokens 根据token推送安卓设备
func (jp *JPushApp) PushAndroidTokens(tokens []string, notification *NotificationAndroid) (result *PushResult, err error) {

	return
}

// PushAndroid 推送安卓设备
func PushAndroid(jp *JPushApp, audience *Audience, notification *NotificationAndroid) (result *PushResult, err error) {
	return jp.PushAndroid(audience, notification)
}

// PushAndroidToken 通过设备token发push
func PushAndroidToken(jp *JPushApp, tokens []string, notification *NotificationAndroid) (result *PushResult, err error) {
	return jp.PushAndroidTokens(tokens, notification)
}
