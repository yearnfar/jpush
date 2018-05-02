package jpush

// Param 参数
type Param struct {
	Platform     interface{}   `json:"platform,omitempty"`
	Audience     interface{}   `json:"audience,omitempty"`
	Notification *Notification `json:"notification,omitempty"`
	Options      []Option      `json:"options,omitempty"`
	Message      *Message      `json:"message,omitempty"`
}

// SetPlatformAll 推送所有平台
func (p *Param) SetPlatformAll() {
	p.Platform = "all"
}

// SetPlatforms 推送指定平台
func (p *Param) SetPlatforms(platforms ...string) {
	p.Platform = platforms
}

// SetAudienceAll 推送所有用户
func (p *Param) SetAudienceAll() {
	p.Audience = "all"
}

// SetAudience 推送指定用户
func (p *Param) SetAudience(audience *Audience) {
	p.Audience = audience
}

// SetNotification 推送android消息
func (p *Param) SetNotification(notification *Notification) {
	p.Notification = notification
}

// AddNotificationAndroid 推送android消息
func (p *Param) AddNotificationAndroid(android *NotificationAndroid) {
	if p.Notification == nil {
		p.Notification = &Notification{Android: android}
	} else {
		p.Notification.Android = android
	}
}

// AddNotificationIos 推送ios消息
func (p *Param) AddNotificationIos(ios *NotificationIos) {
	if p.Notification == nil {
		p.Notification = &Notification{Ios: ios}
	} else {
		p.Notification.Ios = ios
	}
}
