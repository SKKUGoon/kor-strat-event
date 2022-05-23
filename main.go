package main

import (
	"goKorMktEvent/dart"
	"log"
	"time"
)

func main() {
	for {
		p, _ := dart.Dart()
		time.Sleep(10 * time.Second)

		log.Println(p)
	}

}
