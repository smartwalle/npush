package jpush_test

import (
	"gitlab.com/smartwalle/push4go/jpush"
	"testing"
)

func TestJPush_Push(t *testing.T) {
	var p = jpush.PushParam{}
	p.Platform = "all"
	p.Audience = "all"
	p.Notification = &jpush.Notification{}
	p.Notification.Alert = "notification alert"

	p.Notification.Android = &jpush.NotificationAndroid{}
	p.Notification.Android.Title = "title"
	p.Notification.Android.BigText = "big text"
	p.Notification.Android.Style = 1

	p.Notification.IOS = &jpush.NotificationIOS{}
	p.Notification.IOS.Alert = "ios"
	p.Notification.IOS.Badge = 1

	p.Options = &jpush.Options{}
	p.Options.APNSProduction = false

	var rsp, err = client.Push(p)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.MsgId, rsp.SendNo)
}

func TestJPush_GroupPush(t *testing.T) {
	var p = jpush.PushParam{}
	p.Platform = "all"
	p.Audience = "all"
	p.Notification = &jpush.Notification{}
	p.Notification.Alert = "notification alert"

	p.Notification.Android = &jpush.NotificationAndroid{}
	p.Notification.Android.Title = "title"
	p.Notification.Android.BigText = "big text"
	p.Notification.Android.Style = 1

	p.Notification.IOS = &jpush.NotificationIOS{}
	p.Notification.IOS.Alert = "ios"
	p.Notification.IOS.Badge = 1

	p.Options = &jpush.Options{}
	p.Options.APNSProduction = false

	var rsp, err = client.GroupPush("fc9e35a47c3a5e8db61dff89", "e451e479bd3606296926bd04", p)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.Id, rsp.MsgId, rsp.SendNo)
}
