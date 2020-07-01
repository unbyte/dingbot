package dingbot

var defaultBot Bot = New("", "")

// 设置 access_token
func Token(token string) {
	defaultBot.Token(token)
}

// 设置用于安全加密的 secret
func Secret(secret string) {
	defaultBot.Secret(secret)
}

// 测试是否连通，会发送一段文本信息
func Test() error {
	return defaultBot.Text("测试", []string{}, false)
}

// 发送文本信息
func Text(content string, atMobiles []string, atAll bool) error {
	return defaultBot.Text(content, atMobiles, atAll)
}

// 发送链接信息
func Link(title, content, linkURL, picURL string) error {
	return defaultBot.Link(title, content, linkURL, picURL)
}

// 发送 markdown 信息
func Markdown(title, content string, atMobiles []string, atAll bool) error {
	return defaultBot.Markdown(title, content, atMobiles, atAll)
}

// 发送带有按钮的信息
func Action(title, content string, buttons []*Button, btnOrientation BtnOrientation) error {
	return defaultBot.Action(title, content, buttons, btnOrientation)
}

// 发送订阅式卡片信息
func Feed(feed []*FeedItem) error {
	return defaultBot.Feed(feed)
}

// 直接向接口发送自定义信息
func Custom(message []byte) error {
	return defaultBot.Custom(message)
}
