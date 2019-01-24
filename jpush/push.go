package jpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/smartwalle/ngx"
	"net/http"
)

const (
	kPushAPI      = "push"
	kGroupPushAPI = "grouppush"
)

func (this *JPush) Push(param PushParam) (result *PushResponse, err error) {
	if err = this.doRequest(http.MethodPost, kPushAPI, param, &result); err != nil {
		return nil, err
	}
	if result != nil && result.Error != nil {
		return nil, result.Error
	}
	return result, err
}

func (this *JPush) GroupPush(groupKey, groupMasterSecret string, param PushParam) (result *GroupPushResponse, err error) {
	var url = kJPushAPIDomain + kGroupPushAPI

	var req = ngx.NewRequest(http.MethodPost, url)
	req.SetHeader("Authorization", "Basic "+this.Authorization(fmt.Sprintf("group-%s", groupKey), groupMasterSecret))
	req.SetHeader("Content-Type", ngx.K_CONTENT_TYPE_JSON)
	req.SetHeader("Accept", ngx.K_CONTENT_TYPE_JSON)

	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	req.SetBody(bytes.NewReader(data))

	var rMap map[string]map[string]interface{}

	var rsp = req.Exec()
	if err := rsp.UnmarshalJSON(&rMap); err != nil {
		return nil, err
	}

	if rMap == nil {
		return nil, nil
	}

	if rErr, ok := rMap["error"]; ok {
		var code = int(rErr["code"].(float64))
		var msg = rErr["message"].(string)
		return nil, &Error{Code: code, Message: msg}
	}

	result = &GroupPushResponse{}
	for key, value := range rMap {
		result.Id = key
		result.SendNo = value["sendno"].(string)
		result.MsgId = value["msg_id"].(string)
		break
	}

	return result, err
}
