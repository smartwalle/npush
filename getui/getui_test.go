package getui_test

import (
	"gitlab.com/smartwalle/push4go/getui"
	"testing"
)

var client *getui.GeTui

func init() {
	client, _ = getui.New("YAp4wmAzz08KtDNgWLROo4", "I3ctxXhjaD8Oh3RUVf6KBA", "GBSkhrpdw3AAzgh8axV2a4")
}

func TestGeTui_AuthSign(t *testing.T) {
	var rsp, err = client.AuthSign()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.AuthToken)
}
