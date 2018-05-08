package jpush

const (
	PushAPI = "https://api.jpush.cn/v3/push"
)

// JPushApp 极光推送
type JPushApp struct {
	AppKey    string
	AppSecret string

	Client *Client
}

// NewJPushApp 实例化极光推送对象
func NewJPushApp(appKey string, appSecret string) *JPushApp {
	jp := new(JPushApp)
	jp.AppKey = appKey
	jp.AppSecret = appSecret
	jp.Client = NewClient(jp.AppKey, jp.AppSecret)
	return jp
}

// PushAndroid 推送安卓设备
func (jp *JPushApp) PushAndroid(audience *Audience, android *NotificationAndroid) (result *PushResult, err error) {
	param := new(Param)
	param.SetAudience(audience)
	param.SetPlatforms(PlatformAndroid)
	param.AddNotificationAndroid(android)

	if jp.Client == nil {
		jp.Client = NewClient(jp.AppKey, jp.AppSecret)
	}

	result = new(PushResult)
	err = jp.Client.PostJSON(PushAPI, param, result)
	return
}

// PushAndroidTokens 根据token推送安卓设备
func (jp *JPushApp) PushAndroidAlias(alias []string, android *NotificationAndroid) (result *PushResult, err error) {
	param := new(Param)
	param.SetAudience(&Audience{Alias: alias})
	param.SetPlatforms(PlatformAndroid)
	param.AddNotificationAndroid(android)

	if jp.Client == nil {
		jp.Client = NewClient(jp.AppKey, jp.AppSecret)
	}

	result = new(PushResult)
	err = jp.Client.PostJSON(PushAPI, param, result)
	return
}

// PushAndroid 推送安卓设备
func PushAndroid(jp *JPushApp, audience *Audience, msg string) (result *PushResult, err error) {
	return jp.PushAndroid(audience, &NotificationAndroid{Alert: msg})
}

// PushAndroidAlias 通过设备别名发push
func PushAndroidAlias(jp *JPushApp, alias []string, msg string) (result *PushResult, err error) {
	return jp.PushAndroidAlias(alias, &NotificationAndroid{Alert: msg})
}
