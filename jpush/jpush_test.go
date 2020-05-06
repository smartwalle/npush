package jpush_test

import (
	"github.com/smartwalle/push4go/jpush"
	"testing"
)

var client *jpush.Client

func init() {
	client, _ = jpush.New("486eb729aef667639c55c15e", "88670da91fa107d1b1ac52ac")
}

func TestJPush_GetCIdList(t *testing.T) {
	var rsp, err = client.GetCIdList(10, "")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.CIDList)
}
