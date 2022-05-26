package main

import (
	"fmt"
	"goKorMktEvent/watch"
	"log"
	"os"
	"time"
)

const initialArt = `
creator : Sang Il. Bae
github  : https://github.com/SKKUGoon/goKorMktEvent

 /$$   /$$  /$$$$$$  /$$      /$$ /$$      
| $$  /$$/ /$$__  $$| $$  /$ | $$| $$      
| $$ /$$/ | $$  \ $$| $$ /$$$| $$| $$      
| $$$$$/  | $$  | $$| $$/$$ $$ $$| $$      
| $$  $$  | $$  | $$| $$$$_  $$$$| $$      
| $$\  $$ | $$  | $$| $$$/ \  $$$| $$      
| $$ \  $$|  $$$$$$/| $$/   \  $$| $$$$$$$$
|__/  \__/ \______/ |__/     \__/|________/

`

func init() {
	fmt.Print(initialArt)
}

func main() {
	var test = false

	isTest := os.Args
	if len(isTest) != 1 {
		test = true
	}

	// test
	if test {
		reports := []watch.NewReportWatch{watch.TestStruct0, watch.TestStruct1, watch.TestStruct2}
		for _, r := range reports {
			_ = r.Run()
		}
		return
	}

	// not test
	for {
		reports, _ := watch.Rss()

		if len(reports) == 0 {
			//time.Sleep(10 * time.Second)
			log.Println("no info")
		}

		for _, r := range reports {
			_ = r.Run()
		}

		// RSS Feed restriction( 100reqs / min )
		time.Sleep(5 * time.Second)
	}
}
