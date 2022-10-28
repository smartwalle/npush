package jpush

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	kGetCIDListAPI = kJPushAPIDomain + "push/cid"
)

// GetCIdList 获取推送唯一标识符 https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced
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
