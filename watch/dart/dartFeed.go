package dart

import (
	"goKorMktEvent/watch"

	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	sKOSPI  = "유가"
	sKOSDAQ = "코스닥"

	sMusang    = "주요사항보고서(무상증자결정)"
	sYusang    = "주요사항보고서(유상증자결정)"
	sYuMusang  = "주요사항보고서(유무상증자결정)"
	sRedundant = "기재정정"
)

func dartXML(r *http.Response) ([]watch.Content, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("fail to retrieve xml")
	}
	info, err := parseFeed(string(data))
	if err != nil {
		log.Println("fail to parse xml")
		return []watch.Content{}, err
	}

	return info.ChannelDet.Item, nil
}

func parseFeed(s string) (watch.Feed, error) {
	var t watch.Feed
	err := xml.Unmarshal([]byte(s), &t)
	if err != nil {
		log.Fatal(err)
		return t, err
	}

	return t, nil
}

func sortFeed(c []watch.Content) []watch.Content {
	var result []watch.Content

	for _, v := range c {
		_, ok := necessaryFeed(&v)
		if ok {
			result = append(result, v)
		}
	}

	return result
}

func necessaryFeed(c *watch.Content) (string, bool) {
	// Market
	isKOSPI := c.Category == sKOSPI
	isKOSDAQ := c.Category == sKOSDAQ
	if !isKOSPI && !isKOSDAQ {
		return "", false
	}

	// Duplicate
	isDup := strings.Contains(c.Title, sRedundant)
	if isDup {
		return "", false
	}

	// Report Type
	isMusang := strings.Contains(c.Title, sMusang)
	isYusang := strings.Contains(c.Title, sYusang)
	isYuMusang := strings.Contains(c.Title, sYuMusang)
	if isMusang || isYusang || isYuMusang {
		return c.Title, true
	} else {
		return "", false
	}
}

func Rss() ([]watch.Content, error) {
	const url = "https://dart.fss.or.kr/api/todayRSS.xml"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// parse xml
	r, err := dartXML(resp)
	if err != nil {
		return []watch.Content{}, err
	} else {
		return sortFeed(r), nil
	}
}
