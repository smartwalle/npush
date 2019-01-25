package jpush

//////////////////////////////////////////////////////////////////////////////////
// PushParam https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/#_7
type PushParam struct {
	Platform     interface{}   `json:"platform"`               // 必填, 推送平台设置, JPush 当前支持 Android, iOS, Windows Phone 三个平台的推送。其关键字分别为："all", "android", "ios", "winphone"。
	Audience     interface{}   `json:"audience"`               // 必填, 推送设备指定, 推送设备对象，表示一条推送可以被推送到哪些设备列表。确认推送设备对象，JPush 提供了多种方式，比如：别名、标签、注册 ID、分群、广播等。如果要发广播（全部设备），则直接填写 “all”。
	Notification *Notification `json:"notification,omitempty"` // 可选, 通知内容体。是被推送到客户端的内容。与 message 一起二者必须有其一，可以二者并存
	Message      *Message      `json:"message,omitempty"`      // 可选, 消息内容体。是被推送到客户端的内容。与 notification 一起二者必须有其一，可以二者并存
	Options      *Options      `json:"options,omitempty"`      // 可选, 推送参数
	CID          string        `json:"cid,omitempty"`          // 可选, 用于防止 api 调用端重试造成服务端的重复推送而定义的一个标识符。
}

// Audience https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/#audience
type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	Abtest         []string `json:"abtest,omitempty"`
}

// Notification https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/#notification
type Notification struct {
	Alert   string               `json:"alert"`
	Android *NotificationAndroid `json:"android,omitempty"`
	IOS     *NotificationIOS     `json:"ios,omitempty"`
}

type NotificationAndroid struct {
	Alert      string            `json:"alert,omitempty"`        // 必填, 这里指定了，则会覆盖上级统一指定的 alert 信息；内容可以为空字符串，则表示不展示到通知栏。
	Title      string            `json:"title,omitempty"`        // 可选, 如果指定了，则通知里原来展示 App 名称的地方，将展示成这个字段。
	BuilderId  int               `json:"builder_id,omitempty"`   // 可选, 栏样式 ID	Android SDK 可设置通知栏样式，这里根据样式 ID 来指定该使用哪套样式。
	Priority   int               `json:"priority,omitempty"`     // 可选, 默认为 0，范围为 -2～2 ，其他值将会被忽略而采用默认。
	Category   string            `json:"category,omitempty"`     // 可选, 完全依赖 rom 厂商对 category 的处理策略
	Style      int               `json:"style,omitempty"`        // 可选, 默认为 0，还有 1，2，3 可选，用来指定选择哪种通知栏样式，其他值无效。有三种可选分别为 bigText=1，Inbox=2，bigPicture=3。
	AlertType  int               `json:"alert_type,omitempty"`   // 可选, 可选范围为 -1～7 ，对应 Notification.DEFAULT_ALL = -1 或者 Notification.DEFAULT_SOUND = 1, Notification.DEFAULT_VIBRATE = 2, Notification.DEFAULT_LIGHTS = 4 的任意 “or” 组合。默认按照 -1 处理。
	BigText    string            `json:"big_text,omitempty"`     // 可选, 当 style = 1 时可用，内容会被通知栏以大文本的形式展示出来。支持 api 16 以上的 rom。
	Inbox      map[string]string `json:"inbox,omitempty"`        // 可选, 当 style = 2 时可用， json 的每个 key 对应的 value 会被当作文本条目逐条展示。支持 api 16 以上的 rom。
	BigPicPath string            `json:"big_pic_path,omitempty"` // 可选, 当 style = 3 时可用，可以是网络图片 url，或本地图片的 path，目前支持 .jpg 和 .png 后缀的图片。图片内容会被通知栏以大图片的形式展示出来。如果是 http／https 的 url，会自动下载；如果要指定开发者准备的本地图片就填 sdcard 的相对路径。支持 api 16 以上的 rom。
	Extras     map[string]string `json:"extras,omitempty"`       // 可选, 这里自定义 JSON 格式的 Key / Value 信息，以供业务使用。
	LargeIcon  string            `json:"large_icon,omitempty"`   // 可选, 图标路径可以是以http或https开头的网络图片，如：http:jiguang.cn/logo.png ,图标大小不超过 30 k; 也可以是位于drawable资源文件夹的图标路径，如：R.drawable.lg_icon；
	Intent     map[string]string `json:"intent,omitempty"`       // 可选, 使用 intent 里的 url 指定点击通知栏后跳转的目标页面。
}

type NotificationIOS struct {
	Alert            interface{}       `json:"alert,omitempty"`             // 必填, 这里指定内容将会覆盖上级统一指定的 alert 信息；内容为空则不展示到通知栏。支持字符串形式也支持官方定义的 alert payload 结构，在该结构中包含 title 和 subtitle 等官方支持的 key
	Sound            interface{}       `json:"sound,omitempty"`             // 可选, 普通通知： string类型，如果无此字段，则此消息无声音提示；有此字段，如果找到了指定的声音就播放该声音，否则播放默认声音，如果此字段为空字符串，iOS 7 为默认声音，iOS 8 及以上系统为无声音。说明：JPush 官方 SDK 会默认填充声音字段，提供另外的方法关闭声音，详情查看各 SDK 的源码。 告警通知： JSON Object ,支持官方定义的 payload 结构，在该结构中包含 critical 、name 和 volume 等官方支持的 key .
	Badge            int               `json:"badge,omitempty"`             // 可选, 如果不填，表示不改变角标数字，否则把角标数字改为指定的数字；为 0 表示清除。JPush 官方 SDK 会默认填充 badge 值为 "+1",详情参考：badge +1
	ContentAvailable bool              `json:"content-available,omitempty"` // 可选, 推送的时候携带 "content-available":true 说明是 Background Remote Notification，如果不携带此字段则是普通的 Remote Notification。详情参考：Background Remote Notification
	MutableContent   bool              `json:"mutable-content,omitempty"`   // 可选, 推送的时候携带 ”mutable-content":true 说明是支持iOS10的UNNotificationServiceExtension，如果不携带此字段则是普通的 Remote Notification。详情参考：UNNotificationServiceExtension
	Category         string            `json:"category,omitempty"`          // 可选, IOS 8 才支持。设置 APNs payload 中的 "category" 字段值
	Extras           map[string]string `json:"extras,omitempty"`            // 可选, 这里自定义 Key / value 信息，以供业务使用。
	ThreadId         string            `json:"thread-id,omitempty"`         // 可选, ios 的远程通知通过该属性来对通知进行分组，同一个 thread-id 的通知归为一组。
}

// Message https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/#message
type Message struct {
	MsgContent  string            `json:"msg_content"`            // 消息内容本身
	Title       string            `json:"title,omitempty"`        // 消息标题
	ContentType string            `json:"content_type,omitempty"` // 消息内容类型
	Extras      map[string]string `json:"extras,omitempty"`       // JSON 格式的可选参数
}

// Options https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/#options
type Options struct {
	SendNo          int    `json:"sendno,omitempty"`            // 纯粹用来作为 API 调用标识，API 返回时被原样返回，以方便 API 调用方匹配请求与返回。值为 0 表示该 messageid 无 sendno，所以字段取值范围为非 0 的 int.
	TimeToLive      int    `json:"time_to_live,omitempty"`      // 推送当前用户不在线时，为该用户保留多长时间的离线消息，以便其上线时再次推送。默认 86400 （1 天），最长 10 天。设置为 0 表示不保留离线消息，只有推送当前在线的用户可以收到。该字段对 iOS 的 Notification 消息无效。
	OverrideMsgId   int    `json:"override_msg_id,omitempty"`   // 如果当前的推送要覆盖之前的一条推送，这里填写前一条推送的 msg_id 就会产生覆盖效果，即：1）该 msg_id 离线收到的消息是覆盖后的内容；2）即使该 msg_id Android 端用户已经收到，如果通知栏还未清除，则新的消息内容会覆盖之前这条通知；覆盖功能起作用的时限是：1 天。如果在覆盖指定时限内该 msg_id 不存在，则返回 1003 错误，提示不是一次有效的消息覆盖操作，当前的消息不会被推送；该字段仅对 Android 有效。
	APNSProduction  bool   `json:"apns_production"`             // True 表示推送生产环境，False 表示要推送开发环境；如果不指定则为推送生产环境。但注意，JPush 服务端 SDK 默认设置为推送 “开发环境”。该字段仅对 iOS 的 Notification 有效。
	APNSCollapseId  string `json:"apns_collapse_id,omitempty"`  // APNs 新通知如果匹配到当前通知中心有相同 apns-collapse-id 字段的通知，则会用新通知内容来更新它，并使其置于通知中心首位。collapse id 长度不可超过 64 bytes。
	BigPushDuration int    `json:"big_push_duration,omitempty"` // 又名缓慢推送，把原本尽可能快的推送速度，降低下来，给定的 n 分钟内，均匀地向这次推送的目标用户推送。最大值为 1400。未设置则不是定速推送。
}

//////////////////////////////////////////////////////////////////////////////////
type PushResponse struct {
	Error  *Error `json:"error"`
	SendNo string `json:"sendno"`
	MsgId  string `json:"msg_id"`
}

//////////////////////////////////////////////////////////////////////////////////
type GroupPushResponse struct {
	Error  *Error `json:"error"`
	Id     string `json:"id"`
	SendNo string `json:"sendno"`
	MsgId  string `json:"msg_id"`
}
