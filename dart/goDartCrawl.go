package dart

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func dartXML(r *http.Response) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("fail to read xml")
	} else {
		parseXML(string(data))
	}
}

func parseXML(s string) {
	var t ChannelTitle
	err := xml.Unmarshal([]byte(s), &t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
}

func Dart() {
	const url = "https://dart.fss.or.kr/api/todayRSS.xml"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// parse xml
	dartXML(resp)
}
