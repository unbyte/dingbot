package dingbot

import "encoding/json"

type messageType string

const (
	text       messageType = "text"
	link       messageType = "link"
	markdown   messageType = "markdown"
	actionCard messageType = "actionCard"
	feedCard   messageType = "feedCard"
)

func textMessage(content string, atMobiles []string, atAll bool) []byte {
	if atMobiles == nil {
		atMobiles = []string{}
	}
	msg, _ := json.Marshal(map[string]interface{}{
		"msgtype": text,
		"text": map[string]interface{}{
			"content": content,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   atAll,
		},
	})
	return msg
}

func linkMessage(title, content, linkURL, picURL string) []byte {
	msg, _ := json.Marshal(map[string]interface{}{
		"msgtype": link,
		"link": map[string]interface{}{
			"text":       content,
			"title":      title,
			"picURL":     picURL,
			"messageUrl": linkURL,
		},
	})
	return msg
}

func mdMessage(title, content string, atMobiles []string, atAll bool) []byte {
	if atMobiles == nil {
		atMobiles = []string{}
	}
	msg, _ := json.Marshal(map[string]interface{}{
		"msgtype": markdown,
		"markdown": map[string]interface{}{
			"text":  content,
			"title": title,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   atAll,
		},
	})
	return msg
}

type BtnOrientation string

const (
	Vertical   BtnOrientation = "0"
	Horizontal BtnOrientation = "1"
)

type Button struct {
	Title string `json:"title"`
	URL   string `json:"actionURL"`
}

func NewButton(title, url string) *Button {
	return &Button{Title: title, URL: url}
}

func actonMessage(title, content string, buttons []*Button, btnOrientation BtnOrientation) []byte {
	body := map[string]interface{}{
		"msgtype": actionCard,
		"actionCard": map[string]interface{}{
			"text":           content,
			"title":          title,
			"btnOrientation": btnOrientation,
		},
	}
	if len(buttons) == 1 {
		tmp := body["actionCard"].(map[string]interface{})
		tmp["singleTitle"] = buttons[0].Title
		tmp["singleURL"] = buttons[0].URL
	} else if len(buttons) > 1 {
		body["actionCard"].(map[string]interface{})["btns"] = buttons
	}
	msg, _ := json.Marshal(body)
	return msg
}

type FeedItem struct {
	// 消息文本
	Title string `json:"title"`
	// 消息后面的图片的 URL
	MessageURL string `json:"messageURL"`
	// 点击后跳转的链接
	PicURL string `json:"picURL"`
}

func NewFeedItem(title, url, picURL string) *FeedItem {
	return &FeedItem{Title: title, MessageURL: url, PicURL: picURL}
}

func feedMessage(feed []*FeedItem) []byte {
	if feed == nil {
		feed = []*FeedItem{}
	}
	msg, _ := json.Marshal(map[string]interface{}{
		"msgtype": feedCard,
		"feedCard": map[string]interface{}{
			"links": feed,
		},
	})
	return msg
}
