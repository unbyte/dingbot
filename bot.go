package dingbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Bot
type Bot interface {
	// 设置 access_token
	Token(token string)

	// 设置用于安全加密的 secret
	Secret(secret string)

	// 测试是否连通，会发送一段文本信息
	Test() error

	// 发送文本信息
	Text(content string, atMobiles []string, atAll bool) error

	// 发送链接信息
	Link(title, content, linkURL, picURL string) error

	// 发送 markdown 信息
	Markdown(title, content string, atMobiles []string, atAll bool) error

	// 发送带有按钮的信息
	Action(title, content string, buttons []*Button, btnOrientation BtnOrientation) error

	// 发送订阅式卡片信息
	Feed(feed []*FeedItem) error

	// 直接向接口发送自定义信息
	Custom(message []byte) error
}

type bot struct {
	requestURL string
	secret     string

	sign       string
	timestamp  int64 // time.Unix() * 1000
	whenUpdate time.Time

	client *http.Client
}

type response struct {
	ErrCode int    `json:"errcord"`
	ErrMsg  string `json:"errmsg"`
}

func (b *bot) send(msg []byte) error {
	if len(b.secret) == 0 || len(b.requestURL) == 0 {
		return errors.New("发送失败: 未指定 Secret 或 Token")
	}

	if time.Now().After(b.whenUpdate) {
		b.refreshSign()
	}

	realURL := fmt.Sprintf("%s&timestamp=%d&sign=%s", b.requestURL, b.timestamp, b.sign)

	req, err := newRequest(realURL, msg)
	if err != nil {
		return errors.New("构造请求失败: " + err.Error())
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return errors.New("发送请求失败: " + err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("错误的响应: HTTP %d", resp.StatusCode)
	}

	body := readBody(resp.Body)

	var r response
	if err := json.Unmarshal(body, &r); err != nil {
		return errors.New("读取响应失败: " + err.Error())
	}

	if r.ErrCode == 0 {
		return nil
	}

	return errors.New("请求失败: " + r.ErrMsg)
}

func (b *bot) refreshSign() {
	b.whenUpdate = time.Now().Add(40 * time.Minute)
	b.timestamp = time.Now().Unix() * 1e3
	b.sign = sign(b.timestamp, b.secret)
}

// 设置 access_token
func (b *bot) Token(token string) {
	if len(token) == 0 {
		return
	}
	b.requestURL = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", token)
}

// 设置用于安全加密的 secret
func (b *bot) Secret(secret string) {
	b.secret = secret
}

// 测试是否连通，会发送一段文本信息
func (b *bot) Test() error {
	return b.Text("测试", []string{}, false)
}

// 发送文本信息
func (b *bot) Text(content string, atMobiles []string, atAll bool) error {
	return b.send(textMessage(content, atMobiles, atAll))
}

// 发送链接信息
func (b *bot) Link(title, content, linkURL, picURL string) error {
	return b.send(linkMessage(title, content, linkURL, picURL))
}

// 发送 markdown 信息
func (b *bot) Markdown(title, content string, atMobiles []string, atAll bool) error {
	return b.send(mdMessage(title, content, atMobiles, atAll))
}

// 发送带有按钮的信息
func (b *bot) Action(title, content string, buttons []*Button, btnOrientation BtnOrientation) error {
	return b.send(actonMessage(title, content, buttons, btnOrientation))
}

// 发送订阅式卡片信息
func (b *bot) Feed(feed []*FeedItem) error {
	return b.send(feedMessage(feed))
}

// 直接向接口发送自定义信息
func (b *bot) Custom(message []byte) error {
	return b.send(message)
}

var _ Bot = &bot{}

// 新建一个 Bot
func New(token, secret string) Bot {
	b := &bot{
		secret:     secret,
		whenUpdate: time.Now().Add(-1 * time.Minute),
		client:     &http.Client{Timeout: 10 * time.Second},
	}
	b.Token(token)
	return b
}
