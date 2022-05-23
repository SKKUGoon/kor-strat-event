package main

import (
	"goKorMktEvent/watch/dart"
	"log"
	"time"
)

func main() {
	for {
		p, _ := dart.Rss()
		time.Sleep(10 * time.Second)

		log.Println(p)
	}
}
