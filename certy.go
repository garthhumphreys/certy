package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

func main() {
	domainPtr := flag.String("domain", "garthhumphreys.com", "a domain")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("expected 'domain'")
		os.Exit(1)
	}

	if *domainPtr != "" {
		fmt.Println("site:", *domainPtr)
		certSearch(*domainPtr)
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func certSearch(domain string)  {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
		)
	c.OnHTML("table:nth-child(8) > tbody > tr > td > table > tbody > tr > td:nth-child(5)", func(element *colly.HTMLElement) {
		// TODO:
		// - add domains to an array and remove duplicates
		// - print to screen only the unique domains
		fmt.Println(element.Text)
		// fmt.Printf("title:%s \n link:%s \r", element.Text, element.Attr("href"))
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visting", request.URL)
	})

	query := "https://crt.sh/?q=" + domain
	c.Visit(query)
}