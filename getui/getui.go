package getui

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smartwalle/ngx"
	"net/http"
	"strings"
	"time"
)

const (
	kGeTuiAPIDomain = "https://restapi.getui.com/v1/"
)

const (
	kAuthSignAPI = "auth_sign"
)

type GeTui struct {
	appId        string
	appKey       string
	masterSecret string
	token        *Token
}

func New(appId, appKey, masterSecret string) (*GeTui, error) {
	if appId == "" {
		return nil, errors.New("请提供 appId 信息")
	}
	if appKey == "" {
		return nil, errors.New("请提供 appKey 信息")
	}

	if masterSecret == "" {
		return nil, errors.New("请提供 masterSecret 信息")
	}

	var g = &GeTui{}
	g.appId = appId
	g.appKey = appKey
	g.masterSecret = masterSecret
	return g, nil
}

func (this *GeTui) BuildAPI(paths ...string) string {
	var path = fmt.Sprintf("%s%s/", kGeTuiAPIDomain, this.appId)
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			if strings.HasSuffix(path, "/") {
				path = path + p
			} else {
				if strings.HasPrefix(p, "/") {
					path = path + p
				} else {
					path = path + "/" + p
				}
			}
		}
	}
	return path
}

func (this *GeTui) doRequest(method, api string, param interface{}, result interface{}) error {
	if this.token == nil || this.token.IsValid() == false {
		authRsp, err := this.AuthSign()
		if err != nil {
			return err
		}

		if authRsp.IsOk() == false {
			return errors.New(authRsp.Result)
		}

		if authRsp.Token.IsValid() == false {
			return errors.New("auth token is expire")
		}

		this.token = authRsp.Token
	}
	var req = ngx.NewRequest(method, this.BuildAPI(api))
	req.SetHeader("authtoken", this.token.AuthToken)
	req.SetHeader("Content-Type", ngx.K_CONTENT_TYPE_JSON)

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

func (this *GeTui) AuthSign() (result *AuthSignResponse, err error) {
	var req = ngx.NewRequest(http.MethodPost, this.BuildAPI(kAuthSignAPI))
	req.SetHeader("Content-Type", ngx.K_CONTENT_TYPE_JSON)
	req.SetHeader("Accept", ngx.K_CONTENT_TYPE_JSON)

	var now = Microsecond()
	var param = make(map[string]interface{})
	param["sign"] = SignSha256(this.appKey, this.masterSecret, now)
	param["timestamp"] = now
	param["appkey"] = this.appKey

	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	req.SetBody(bytes.NewReader(data))

	var rsp = req.Exec()
	if err := rsp.UnmarshalJSON(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func SignSha256(appKey, masterSecret string, microsecond int64) (sign string) {
	sha256Ctx := sha256.New()
	sha256Ctx.Write([]byte(fmt.Sprintf("%s%d%s", appKey, microsecond, masterSecret)))
	cipherStr := sha256Ctx.Sum(nil)
	sign = strings.ToUpper(hex.EncodeToString(cipherStr))
	return sign
}

func Microsecond() int64 {
	var now = time.Now()
	return now.UnixNano() / 1000000
}
