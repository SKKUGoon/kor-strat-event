package watch

import (
	"encoding/xml"
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

func sortFeed(c []Content) []NewReportWatch {
	var result []NewReportWatch

	for _, v := range c {
		var report NewReportWatch
		ok, typeNum := necessaryFeed(&v)

		if ok {
			report.Ctx = v
			report.T = typeNum
			result = append(result, report)
		}
	}

	return result
}

func necessaryFeed(c *Content) (bool, int) {
	/*
		RSS 피드에서 받은 정보가 필요한 이벤트에 대한 정보인지 구분함.
		  1: 무상증자
		  2: 유상증자
		  3: 유무상증자
		  4: 전환사채
		 -1: 필요한 정보 아님
		-----
		Identifies whether the information gathered in the RSS feed is necessary
		  1: Bonus Issue
		  2: Rights Issue
		  3: Bonus & Right Issue
		  4: Convertible Bond Issue
		 -1: Not Necessary
	*/

	// Market
	isKOSPI := c.Category == sKOSPI
	isKOSDAQ := c.Category == sKOSDAQ
	if !isKOSPI && !isKOSDAQ {
		return false, -1
	}

	// Duplicate
	isDup := strings.Contains(c.Title, sRedundant)
	if isDup {
		return false, -1
	}

	// Report Type
	isMusang := strings.Contains(c.Title, sMusang)
	isYusang := strings.Contains(c.Title, sYusang)
	isYuMusang := strings.Contains(c.Title, sYuMusang)
	isConvert := strings.Contains(c.Title, sConvert)

	switch {
	case isMusang:
		return true, 1
	case isYusang:
		return true, 2
	case isYuMusang:
		return true, 3
	case isConvert:
		return true, 4
	default:
		return false, -1
	}
}

func Rss() ([]NewReportWatch, error) {
	const url = "https://dart.fss.or.kr/api/todayRSS.xml"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// parse xml
	r, err := dartXML(resp)
	if err != nil {
		return []NewReportWatch{}, err
	} else {
		return sortFeed(r), nil
	}
}
