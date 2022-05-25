package main

import (
	"goKorMktEvent/watch"
	"log"
)

func main() {
	for {
		reports, _ := watch.Rss()
		if len(reports) == 0 {
			//time.Sleep(10 * time.Second)
			log.Println("no info")
			break
		}

		for _, r := range reports {
			log.Println(r.Ctx.Title)
			r.Run()
		}
		break
	}
}
