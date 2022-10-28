package jpush

type CIDListResponse struct {
	Error   *Error   `json:"error"`
	CIDList []string `json:"cidlist"`
}
