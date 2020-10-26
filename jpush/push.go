package jpush

import (
	"fmt"
	"github.com/smartwalle/ngx"
	"net/http"
)

const (
	kPushAPI         = kJPushAPIDomain + "push"
	kGroupPushAPI    = kJPushAPIDomain + "grouppush"
	kPushValidateAPI = kJPushAPIDomain + "push/validate"
)

// Push 推送 https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/#_1
func (this *Client) Push(param PushParam) (result *PushResponse, err error) {
	if err = this.doRequest(http.MethodPost, kPushAPI, param, &result); err != nil {
		return nil, err
	}
	if result != nil && result.Error != nil {
		return nil, result.Error
	}
	return result, err
}

// GroupPush 应用分组推送 https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/#group-push-api
func (this *Client) GroupPush(groupKey, groupMasterSecret string, param PushParam) (result *GroupPushResponse, err error) {
	var req = ngx.NewRequest(http.MethodPost, kGroupPushAPI)
	req.SetHeader("Authorization", "Basic "+this.Authorization(fmt.Sprintf("group-%s", groupKey), groupMasterSecret))
	req.SetHeader("Accept", ngx.ContentTypeJSON)
	req.WriteJSON(param)

	var rMap map[string]interface{}

	var rsp = req.Exec()
	fmt.Println(rsp.MustString())
	if err := rsp.UnmarshalJSON(&rMap); err != nil {
		return nil, err
	}

	if rMap == nil {
		return nil, nil
	}

	if rErr, ok := rMap["error"]; ok {
		var rMap = rErr.(map[string]interface{})
		var code int
		if v, ok := rMap["code"]; ok {
			code = int(v.(float64))
		}

		var msg string
		if v, ok := rMap["message"]; ok {
			msg = v.(string)
		}
		return nil, &Error{Code: code, Message: msg}
	}

	result = &GroupPushResponse{}
	for key, value := range rMap {
		if key == "group_msgid" {
			result.GroupMsgId = value.(string)
		} else {
			if vMap, ok := value.(map[string]interface{}); ok {
				var r = &GroupPushResult{}
				r.Id = key
				r.SendNo = vMap["sendno"].(string)
				r.MsgId = vMap["msg_id"].(string)
				result.Result = append(result.Result, r)
			}
		}
	}

	return result, err
}

// PushValidate 推送校验 API https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/#api
func (this *Client) PushValidate(param PushParam) (result *PushResponse, err error) {
	if err = this.doRequest(http.MethodPost, kPushValidateAPI, param, &result); err != nil {
		return nil, err
	}
	if result != nil && result.Error != nil {
		return nil, result.Error
	}
	return result, err
}
