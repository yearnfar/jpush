package jpush

// Audience 推送目标
type Audience struct {
	// 数组。多个标签之间是 OR 的关系，即取并集。
	Tag []string `json:"tag,omitempty"`

	// 数组。多个标签之间是 AND 关系，即取交集。
	TagAnd []string `json:"tag_and,omitempty"`

	// 数组。多个标签之间，先取多标签的并集，再对该结果取补集。
	TagNot []string `json:"tag_not,omitempty"`

	// Alias 数组。多个别名之间是 OR 关系，即取并集。
	Alias string `json:"alias,omitempty"`

	// RegistrationID 数组。多个注册ID之间是 OR 关系，即取并集。
	RegistrationID []string `json:"registration_id,omitempty"`

	// Segment 在页面创建的用户分群的 ID。定义为数组，但目前限制一次只能推送一个。
	Segment []string `json:"segment,omitempty"`
}
