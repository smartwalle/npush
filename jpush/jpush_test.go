package jpush_test

import (
	"github.com/smartwalle/push4go/jpush"
	"testing"
)

var client *jpush.Client

func init() {
	client, _ = jpush.New("4bd172536c75345d7243c7a0", "219a7aeeb09b362ac350a17d")
}

func TestJPush_GetCIdList(t *testing.T) {
	var rsp, err = client.GetCIdList(10, "")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.CIDList)
}
