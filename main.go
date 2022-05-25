package main

import (
	"goKorMktEvent/watch"
	"log"
	"time"
)

func main() {
	for {
		reports, _ := watch.Rss()
		if len(reports) == 0 {
			//time.Sleep(10 * time.Second)
			log.Println("no info")
		}

		for _, r := range reports {
			log.Println(r.Ctx.Title)
			r.Run()
		}

		// RSS Feed restriction( 100reqs / min )
		time.Sleep(5 * time.Second)
	}
}
