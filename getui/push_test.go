package getui_test

import (
	"fmt"
	"gitlab.com/smartwalle/push4go/getui"
	"testing"
)

func TestGeTui_Push(t *testing.T) {
	var p = getui.PushParam{}
	var rsp, err = client.Push(p)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rsp.Result)
}
