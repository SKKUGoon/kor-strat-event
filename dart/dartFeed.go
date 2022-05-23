package dart

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func dartXML(r *http.Response) ([]Content, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("fail to retrieve xml")
	}
	info, err := parseFeed(string(data))
	if err != nil {
		log.Println("fail to parse xml")
		return []Content{}, err
	}

	return info.ChannelDet.Item, nil
}

func parseFeed(s string) (Feed, error) {
	var t Feed
	err := xml.Unmarshal([]byte(s), &t)
	if err != nil {
		log.Fatal(err)
		return t, err
	}

	return t, nil
}

func sortFeed(c *Content) (string, bool) {
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

func Dart() (Content, error) {
	const url = "https://dart.fss.or.kr/api/todayRSS.xml"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// parse xml
	r, err := dartXML(resp)
	if err != nil {
		return Content{}, err
	} else {
		for _, v := range r {
			title, ok := sortFeed(&v)
			if ok {
				fmt.Println(title)
				fmt.Println(v.Link)
			}
		}
	}
	return Content{}, err
}
