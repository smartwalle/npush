package jpush

import (
	"encoding/base64"
	"errors"
	"github.com/smartwalle/ngx"
)

const (
	kJPushAPIDomain = "https://api.jpush.cn/v3/"
)

type Client struct {
	appKey        string
	masterSecret  string
	authorization string
}

func New(appKey, masterSecret string) (*Client, error) {
	if appKey == "" {
		return nil, errors.New("请提供 appKey 信息")
	}

	if masterSecret == "" {
		return nil, errors.New("请提供 masterSecret 信息")
	}

	var p = &Client{}
	p.appKey = appKey
	p.masterSecret = masterSecret
	p.authorization = p.Authorization(p.appKey, p.masterSecret)
	return p, nil
}

func (this *Client) Authorization(key, secret string) string {
	return base64.StdEncoding.EncodeToString([]byte(key + ":" + secret))
}

func (this *Client) doRequest(method, url string, param interface{}, result interface{}) error {
	var req = ngx.NewRequest(method, url)
	req.SetHeader("Authorization", "Basic "+this.authorization)
	req.SetHeader("Accept", ngx.ContentTypeJSON)
	req.WriteJSON(param)

	var rsp = req.Exec()
	if err := rsp.UnmarshalJSON(&result); err != nil {
		return err
	}

	return nil
}
