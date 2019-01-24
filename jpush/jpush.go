package jpush

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smartwalle/ngx"
	"net/http"
	"net/url"
)

const (
	kJPushAPIDomain = "https://api.jpush.cn/v3/"
)

const (
	kGetCIDListAPI = "push/cid"
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
	p.authorization = p.Authorization(p.appKey, p.masterSecret)
	return p, nil
}

func (this *JPush) Authorization(key, secret string) string {
	return base64.StdEncoding.EncodeToString([]byte(key + ":" + secret))
}

func (this *JPush) doRequest(method, api string, param interface{}, result interface{}) error {
	var url = kJPushAPIDomain + api

	var req = ngx.NewRequest(method, url)
	req.SetHeader("Authorization", "Basic "+this.authorization)
	req.SetHeader("Content-Type", ngx.K_CONTENT_TYPE_JSON)
	req.SetHeader("Accept", ngx.K_CONTENT_TYPE_JSON)

	data, err := json.Marshal(param)
	if err != nil {
		return err
	}
	req.SetBody(bytes.NewReader(data))

	var rsp = req.Exec()
	if err := rsp.UnmarshalJSON(&result); err != nil {
		return err
	}

	return nil
}

func (this *JPush) GetCIdList(count int, cType string) (result *CIDListResponse, err error) {
	if count <= 0 {
		count = 10
	}

	var v = url.Values{}
	v.Set("count", fmt.Sprintf("%d", count))
	if cType != "" {
		v.Set("type", cType)
	}

	var api = fmt.Sprintf("%s?%s", kGetCIDListAPI, v.Encode())

	if err = this.doRequest(http.MethodGet, api, nil, &result); err != nil {
		return nil, err
	}
	if result != nil && result.Error != nil {
		return nil, result.Error
	}
	return result, err
}
