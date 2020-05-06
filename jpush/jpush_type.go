package jpush

import (
	"fmt"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (this *Error) Error() string {
	return fmt.Sprintf("%d-%s", this.Code, this.Message)
}

type CIDListResponse struct {
	Error   *Error   `json:"error"`
	CIDList []string `json:"cidlist"`
}
