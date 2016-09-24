package huaweipush

import "testing"

func init() {
	Init("clientID", "clientSecret")
}

var defaultClient *HuaweiPushClient

func Init(clientID, clientSecret string) {
	defaultClient = &HuaweiPushClient{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func TestHuaweiPushClient_SingleSend(t *testing.T) {
	result, err := defaultClient.SingleSend(NewSingleNotification("deviceToken", "hi baby").SetRequestID("requestID"))
	if err != nil {
		t.Errorf("err=%v\n", err)
		return
	}
	result, err = defaultClient.SingleSend(NewSingleNotification("deviceToken", "hi baby").SetRequestID("requestID"))
	if err != nil {
		t.Errorf("err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_NotificationSend(t *testing.T) {
	result, err := defaultClient.NotificationSend(
		NewNotification(1, 1).
			addTokens("tokens1").
			addTokens("tokens2").
			setAndroid(
			NewAndroidMessage("hi", "baby")))

	if err != nil {
		t.Errorf("err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_SetUserTag(t *testing.T) {
	result, err := defaultClient.SetUserTag("token1", "地区", "上海")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_SetUserTag err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_QueryAppTags(t *testing.T) {
	result, err := defaultClient.QueryAppTags()
	if err != nil {
		t.Errorf("TestHuaweiPushClient_QueryAppTags err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_DeleteUserTag(t *testing.T) {
	result, err := defaultClient.DeleteUserTag("token1", "地区")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_DeleteUserTag err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_QueryUserTag(t *testing.T) {
	result, err := defaultClient.QueryUserTag("token1")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_QueryUserTag err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_QueryMsgResult(t *testing.T) {
	result, err := defaultClient.QueryMsgResult("1474696910794906661310000000", "token1")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_QueryUserTag err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_GetTokenByDate(t *testing.T) {
	result, err := defaultClient.GetTokenByDate("2016-09-23")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_GetTokenByDate err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}
