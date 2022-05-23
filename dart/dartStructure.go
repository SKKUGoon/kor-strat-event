package dart

import "encoding/xml"

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

const (
	sKOSPI  = "유가"
	sKOSDAQ = "코스닥"

	sMusang    = "주요사항보고서(무상증자결정)"
	sYusang    = "주요사항보고서(유상증자결정)"
	sYuMusang  = "주요사항보고서(유무상증자결정)"
	sRedundant = "기재정정"
)
