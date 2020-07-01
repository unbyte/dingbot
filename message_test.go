package dingbot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextMessage(t *testing.T) {
	a := assert.New(t)
	expected1 := []byte(`{"at":{"atMobiles":[],"isAtAll":false},"msgtype":"text","text":{"content":"test"}}`)
	result1 := textMessage("test", nil, false)
	result2 := textMessage("test", []string{}, false)
	a.Equal(expected1, result1)
	a.Equal(result1, result2)

	expected2 := []byte(`{"at":{"atMobiles":["156xxxx8827","189xxxx8325"],"isAtAll":false},"msgtype":"text","text":{"content":"test"}}`)
	result3 := textMessage("test", []string{"156xxxx8827", "189xxxx8325"}, false)
	a.Equal(expected2, result3)
}

func TestLinkMessage(t *testing.T) {
	a := assert.New(t)
	expected := []byte(`{"link":{"messageUrl":"linkURL","picURL":"picURL","text":"test","title":"test"},"msgtype":"link"}`)
	result := linkMessage("test", "test", "linkURL", "picURL")
	a.Equal(expected, result)
}

func TestMarkdownMessage(t *testing.T) {
	a := assert.New(t)
	expected := []byte(`{"at":{"atMobiles":[],"isAtAll":false},"markdown":{"text":"# test","title":"test"},"msgtype":"markdown"}`)
	result := mdMessage("test", "# test", nil, false)
	a.Equal(expected, result)
}

func TestActionMessage(t *testing.T) {
	a := assert.New(t)
	type testCase struct {
		Expected []byte
		Buttons  []*Button
	}
	testCases := []*testCase{
		{Expected: []byte(`{"actionCard":{"btnOrientation":"0","text":"test","title":"test"},"msgtype":"actionCard"}`), Buttons: nil},
		{Expected: []byte(`{"actionCard":{"btnOrientation":"0","singleTitle":"btnTitle","singleURL":"btnURL","text":"test","title":"test"},"msgtype":"actionCard"}`),
			Buttons: []*Button{NewButton("btnTitle", "btnURL")}},
		{Expected: []byte(`{"actionCard":{"btnOrientation":"0","btns":[{"title":"btnTitleA","actionURL":"btnURLA"},{"title":"btnTitleB","actionURL":"btnURLB"}],"text":"test","title":"test"},"msgtype":"actionCard"}`),
			Buttons: []*Button{NewButton("btnTitleA", "btnURLA"), NewButton("btnTitleB", "btnURLB")}},
	}
	for _, c := range testCases {
		a.Equal(c.Expected, actonMessage("test", "test", c.Buttons, Vertical))
	}
}

func TestFeedMessage(t *testing.T) {
	a := assert.New(t)
	type testCase struct {
		Expected  []byte
		FeedItems []*FeedItem
	}
	testCases := []*testCase{
		{Expected: []byte(`{"feedCard":{"links":[]},"msgtype":"feedCard"}`), FeedItems: nil},
		{Expected: []byte(`{"feedCard":{"links":[{"title":"feedTitle","messageURL":"feedURL","picURL":"feedPic"}]},"msgtype":"feedCard"}`),
			FeedItems: []*FeedItem{NewFeedItem("feedTitle", "feedURL", "feedPic")}},
		{Expected: []byte(`{"feedCard":{"links":[{"title":"feedTitleA","messageURL":"feedURLA","picURL":"feedPicA"},{"title":"feedTitleB","messageURL":"feedURLB","picURL":"feedPicB"}]},"msgtype":"feedCard"}`),
			FeedItems: []*FeedItem{NewFeedItem("feedTitleA", "feedURLA", "feedPicA"),
				NewFeedItem("feedTitleB", "feedURLB", "feedPicB")}},
	}
	for _, c := range testCases {
		a.Equal(c.Expected, feedMessage(c.FeedItems))
	}
}
