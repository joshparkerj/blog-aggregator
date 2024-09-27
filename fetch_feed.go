package main

import (
	"bytes"
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedUrl string) (feed *RSSFeed, err error) {
	reader := bytes.NewReader(make([]byte, 0))
	req, err := http.NewRequest("GET", feedUrl, reader)
	if err != nil {
		return
	}

	req.Header.Add("User-Agent", "gator")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	feed = &RSSFeed{}
	err = xml.Unmarshal(dat, feed)
	if err != nil {
		return
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return
}
