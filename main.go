package main

import (
	"goKorMktEvent/watch/dart"
	"goKorMktEvent/watch/parsing"
	"log"
	"time"
)

func main() {
	for {
		reports, _ := dart.Rss()
		if len(reports) == 0 {
			//time.Sleep(10 * time.Second)
			log.Println("no info")
			break
		} else {
			log.Println(reports)
		}
		time.Sleep(10 * time.Second)
	}
	parsing.RawReportText("http://dart.fss.or.kr/dsaf001/main.do?rcpNo=20220523000062")
}
