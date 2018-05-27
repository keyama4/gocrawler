package gocrawler

import (
	"github.com/sclevine/agouti"
	log "github.com/sirupsen/logrus"
)

// CrawlOnChrome is crawling with common process on chrome driver
func CrawlOnChrome(callback func(page *agouti.Page)) {
	crawl(callback, agouti.ChromeDriver())
}

// CrawlOnPhantom is crawling with common process on PhantomJS
func CrawlOnPhantom(callback func(page *agouti.Page)) {
	crawl(callback, agouti.PhantomJS())
}

func crawl(callback func(page *agouti.Page), driver *agouti.WebDriver) {
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()
	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	callback(page)
}
