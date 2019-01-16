package jpush

import (
	"encoding/base64"
	"errors"
	"github.com/smartwalle/ngx"
)

const (
	kJPushAPIDomain = "https://api.jpush.cn/v3/"
)

type JPush struct {
	appKey        string
	masterSecret  string
	authorization string
}

func New(appKey, masterSecret string) (*JPush, error) {
	if appKey == "" {
		return nil, errors.New("请提供 appKey 信息")
	}

	if masterSecret == "" {
		return nil, errors.New("请提供 masterSecret 信息")
	}

	var p = &JPush{}
	p.appKey = appKey
	p.masterSecret = masterSecret
	p.authorization = p.Authorization()
	return p, nil
}

func (this *JPush) Authorization() string {
	return base64.StdEncoding.EncodeToString([]byte(this.appKey + ":" + this.masterSecret))
}

func (this *JPush) doRequest(api string, param interface{}, result interface{}) error {
	var url = kJPushAPIDomain + api

	var req = ngx.NewRequest("POST", url)
	req.SetHeader("Authorization", "Basic "+this.authorization)
	req.SetHeader("Content-Type", ngx.K_CONTENT_TYPE_TEXT)
	req.SetHeader("Accept", ngx.K_CONTENT_TYPE_JSON)
	if err := req.MarshalJSON(param); err != nil {
		return err
	}

	var rsp = req.Exec()
	if err := rsp.UnmarshalJSON(&result); err != nil {
		return err
	}

	return nil
}
