package getui

//////////////////////////////////////////////////////////////////////////////////
// PushParam http://docs.getui.com/getui/server/rest/push/
type PushParam struct {
	Message      *Message      `json:"message"`
	Notification *Notification `json:"notification,omitempty"`
	PushInfo     *PushInfo     `json:"push_info,omitempty"`
	CID          string        `json:"cid,omitempty"`
	RequestId    string        `json:"requestid"`
	Alias        string        `json:"alias,omitempty"`
}

type Message struct {
	AppKey            string `json:"appkey"`                        // 注册应用时生成的appkey
	IsOffline         bool   `json:"is_offline,omitempty"`          // 是否离线推送 可选 默认true
	OfflineExpireTime int    `json:"offline_expire_time,omitempty"` // 消息离线存储有效期，单位：ms 默认24小时
	PushNetWorkType   int    `json:"push_network_type,omitempty"`   // 选择推送消息使用网络类型，0：不限制，1：wifi 默认0
	MsgType           string `json:"msgtype"`                       // 消息应用类型，可选项：notification、link、notypopload、transmission
}

type Notification struct {
	TransmissionType    bool   `json:"transmission_type"`    // 否 收到消息是否立即启动应用，true为立即启动，false则广播等待启动，默认是否
	TransmissionContent string `json:"transmission_content"` // 否 透传内容
	DurationBegin       string `json:"duration_begin"`       // 否 设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss
	DurationEnd         string `json:"duration_end"`         // 否 设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss
	Style               *Style `json:"style"`                // 是 通知栏消息布局样式，见底下Style说明
}

type Link struct {
	URL           string `json:"url,omitempty"`            // 打开网址 可选
	DurationBegin string `json:"duration_begin,omitempty"` // 设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd   string `json:"duration_end,omitempty"`   // 设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	Style         *Style `json:"style,omitempty"`          // 通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
}

type Style struct {
	Type        int    `json:"type,omitempty"`
	Text        string `json:"text,omitempty"`
	Title       string `json:"title,omitempty"`
	Logo        string `json:"logo,omitempty"`
	LogoURL     string `json:"logourl,omitempty"`
	IsRing      bool   `json:"is_ring"`
	IsVibrate   bool   `json:"is_vibrate"`
	IsClearable bool   `json:"is_clearable"`
}

type PushInfo struct {
	Aps struct {
		Alert struct {
			Title string `json:"title,omitempty"`
			Body  string `json:"body,omitempty"`
		} `json:"alert"`
		AutoBadge        string `json:"autoBadge,omitempty"`
		ContentAvailable int    `json:"content-available,omitempty"`
	} `json:"aps"`
	Transmission string `json:"transmission,omitempty"`
}

//////////////////////////////////////////////////////////////////////////////////
type PushResponse struct {
	*Response
	TaskId string `json:"task_id"`
	Status string `json:"status"`
}
