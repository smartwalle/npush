package aps

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

const (
	kDevelopmentURL     = "https://api.development.push.apple.com"
	kDevelopmentURL2197 = "https://api.development.push.apple.com:2197"
	kProductionURL      = "https://api.push.apple.com"
	kProductionURL2197  = "https://api.push.apple.com:2197"
)

const (
	kMaxPayload = 4096
)

type Header struct {
	ID          string
	CollapseID  string
	Expiration  time.Time
	LowPriority bool
	Topic       string
}

func (this *Header) set(reqHeader http.Header) {
	if this.ID != "" {
		reqHeader.Set("apns-id", this.ID)
	}

	if this.CollapseID != "" {
		reqHeader.Set("apns-collapse-id", this.CollapseID)
	}

	if !this.Expiration.IsZero() {
		reqHeader.Set("apns-expiration", strconv.FormatInt(this.Expiration.Unix(), 10))
	}

	if this.LowPriority {
		reqHeader.Set("apns-priority", "5")
	}

	if this.Topic != "" {
		reqHeader.Set("apns-topic", this.Topic)
	}
}

type Payload struct {
	Alert            Alert
	Badge            uint
	Sound            string
	ContentAvailable bool
	Category         string
	MutableContent   bool
	ThreadID         string
	userInfo         map[string]interface{}
}

func (this *Payload) AddUserInfo(key string, value interface{}) {
	if key == "" || value == nil {
		return
	}
	if this.userInfo == nil {
		this.userInfo = make(map[string]interface{})
	}
	this.userInfo[key] = value
}

type Alert struct {
	Title        string   `json:"title,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	Subtitle     string   `json:"subtitle,omitempty"`
	Body         string   `json:"body,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
}

func (this *Alert) isSimple() bool {
	return len(this.Title) == 0 && len(this.Subtitle) == 0 &&
		len(this.LaunchImage) == 0 &&
		len(this.TitleLocKey) == 0 && len(this.TitleLocArgs) == 0 &&
		len(this.LocKey) == 0 && len(this.LocArgs) == 0 && len(this.ActionLocKey) == 0
}

func (this *Alert) isEmpty() bool {
	return len(this.Body) == 0 && this.isSimple()
}

func (this *Payload) toMap() map[string]interface{} {
	aps := make(map[string]interface{}, 5)

	if !this.Alert.isEmpty() {
		if this.Alert.isSimple() {
			aps["alert"] = this.Alert.Body
		} else {
			aps["alert"] = this.Alert
		}
	}
	aps["badge"] = this.Badge
	if this.Sound != "" {
		aps["sound"] = this.Sound
	}
	if this.ContentAvailable {
		aps["content-available"] = 1
	}
	if this.Category != "" {
		aps["category"] = this.Category
	}
	if this.MutableContent {
		aps["mutable-content"] = 1
	}
	if this.ThreadID != "" {
		aps["thread-id"] = this.ThreadID
	}
	if this.userInfo != nil {
		aps["userinfo"] = this.userInfo
	}
	return map[string]interface{}{"aps": aps}
}

func (this Payload) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.toMap())
}

type PushResponse struct {
	Reason string `json:"reason"`
}
