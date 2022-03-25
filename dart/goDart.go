package dart

type Content struct {
	Title    string `xml:"title"`
	Link     string `xml:"link"`
	Category string `xml:"category"`
	PubDate  string `xml:"pubDate"`
	Creator  string `xml:"dc:creator"`
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
