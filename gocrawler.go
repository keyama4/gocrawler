package gocrawler

import (
	"github.com/sclevine/agouti"
	log "github.com/sirupsen/logrus"
)

// CrawlOnChrome is crawling with common process on chrome driver
func CrawlOnChrome(callbacks ...func(page *agouti.Page)) {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()
	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	for _, callback := range callbacks {
		callback(page)
	}
}

// CrawlOnPhantom is crawling with common process on PhantomJS
func CrawlOnPhantom(callbacks ...func(page *agouti.Page)) {
	driver := agouti.PhantomJS()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()
	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	for _, callback := range callbacks {
		callback(page)
	}
}
