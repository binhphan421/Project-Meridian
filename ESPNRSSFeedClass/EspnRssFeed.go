package ESPNRSSFeedClass

import (
	. "Meridian/ArticleClass"
)

type Image struct {
	Url    string `xml:"url"`
	Title  string `xml:"title"`
	Link   string `xml:"link"`
	Width  int    `xml:"width"`
	Height int    `xml:"height"`
}

type Channel struct {
	Title         string    `xml:"title"`
	Desc          string    `xml:"description"`
	Link          string    `xml:"link"`
	TTL           string    `xml:"ttl"`
	Language      string    `xml:"language"`
	Generator     string    `xml:"generator"`
	Copyright     string    `xml:"copyright"`
	LastBuildDate string    `xml:"lastbuilddate"`
	Images        Image     `xml:"image"`
	Articles      []Article `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}
