package getui

import "net/http"

const (
	kPushAPI = "push_single"
)

// Push http://docs.getui.com/getui/server/rest/push/
func (this *GeTui) Push(param PushParam) (result *PushResponse, err error) {
	if err = this.doRequest(http.MethodPost, kPushAPI, param, &result); err != nil {
		return nil, err
	}
	return result, err
}
