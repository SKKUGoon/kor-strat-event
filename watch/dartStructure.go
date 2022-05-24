package watch

import (
	"encoding/xml"
	"time"
)

type Feed struct {
	ChannelInit xml.Name     `xml:"rss"`
	ChannelDet  ChannelTitle `xml:"channel"`
}

type ChannelTitle struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Desc      string    `xml:"description"`
	Language  string    `xml:"language"`
	Copyright string    `xml:"copyright"`
	PubDate   string    `xml:"pubDate"`
	Item      []Content `xml:"item"`
}

type Content struct {
	Title    string `xml:"title"`
	Link     string `xml:"link"`
	Category string `xml:"category"`
	PubDate  string `xml:"pubDate"`
	Creator  string `xml:"dc:creator"`
}

type ReportMap map[string]string

type BonusIssue struct {
	Company    string
	StkPrice   int64
	NewStock   float32
	ReportDate time.Time
	LockDate   time.Time
}

type RightsIssue struct {
	Company  string
	NewStock float32
}
