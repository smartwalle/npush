package getui

import "strconv"

type Response struct {
	Result string `json:"result"`
}

func (this *Response) IsOk() bool {
	if this.Result == "ok" {
		return true
	}
	return false
}

type AuthSignResponse struct {
	*Response
	*Token
}

type Token struct {
	ExpireTime string `json:"expire_time"`
	AuthToken  string `json:"auth_token"`
}

func (this *Token) IsValid() bool {
	t, err := strconv.ParseInt(this.ExpireTime, 10, 64)
	if err != nil {
		return false
	}
	var now = Microsecond()
	if t > now+1000 {
		return true
	}
	return false
}
