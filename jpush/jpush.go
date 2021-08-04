package jpush

import (
	"encoding/base64"
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
	kGetCIDListAPI = kJPushAPIDomain + "push/cid"
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

// GetCIdList 获取推送唯一标识符 https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/#cid
func (this *Client) GetCIdList(count int, cType string) (result *CIDListResponse, err error) {
	if count <= 0 {
		count = 10
	}

	var v = url.Values{}
	v.Set("count", fmt.Sprintf("%d", count))
	if cType != "" {
		v.Set("type", cType)
	}

	var url = fmt.Sprintf("%s?%s", kGetCIDListAPI, v.Encode())

	if err = this.doRequest(http.MethodGet, url, nil, &result); err != nil {
		return nil, err
	}
	if result != nil && result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}
