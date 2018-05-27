package gocrawler

import (
	"github.com/sclevine/agouti"
	log "github.com/sirupsen/logrus"
)

// CrawlOnChrome is crawling with common process on chrome driver
func CrawlOnChrome(onCrawl func(page *agouti.Page)) {
	crawl(agouti.ChromeDriver(), onCrawl)
}

// CrawlOnPhantom is crawling with common process on PhantomJS
func CrawlOnPhantom(onCrawl func(page *agouti.Page)) {
	crawl(agouti.PhantomJS(), onCrawl)
}

func crawl(driver *agouti.WebDriver, onCrawl func(page *agouti.Page)) {
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()
	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	onCrawl(page)
}
