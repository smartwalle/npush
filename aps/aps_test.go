package aps_test

import (
	"github.com/smartwalle/npush/aps"
	"testing"
)

func TestNew(t *testing.T) {
	client, err := aps.New("com.hoteldelins.protals", "./dis.p12", "123456", true)
	if err != nil {
		t.Fatal(err)
	}

	var p = &aps.APS{}
	p.Alert.Title = "title"
	p.Alert.Body = "body"
	p.Badge = 0
	p.AddUserInfo("sss", "haha")

	rsp, err := client.Push("a720efde85cc2183d186b0366e77281987c53eb0b069717b7d2f6c3bc0d4f0b7", nil, p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}
