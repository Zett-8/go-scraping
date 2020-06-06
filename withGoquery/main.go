package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://github.com/Zett-8")
	if err != nil {
		panic(err)
	}

	// scrape only 'a' tag
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Println(url)
	})

	fmt.Println("-----------------------------------------------------")

	// or able to choose tags like jQuery
	doc.Find("a.text-bold").Each(func(_ int, s *goquery.Selection) {
		repo, _ := s.Attr("href")
		fmt.Println(repo)
	})
}
