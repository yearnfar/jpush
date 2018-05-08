package jpush

import (
	"testing"
	"time"
)

func TestPushAndroid(t *testing.T) {
	jp := &JPushApp{
		AppKey:    "yourAppKey",
		AppSecret: "yourAppSecret",
	}

	audience := &Audience{
		Alias: []string{"1111", "2222"},
	}

	result, err := PushAndroid(jp, audience, "hello world! ")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(result)
}

func TestPushAndroidAlias(t *testing.T) {
	jp := NewJPushApp("yourAppKey", "yourAppSecret")

	result, err := PushAndroidAlias(jp, []string{"1111", "2222"}, "hello world!")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(result)
}
