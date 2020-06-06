package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	o := "olive-gallery.dev.s3-website-ap-northeast-1.amazonaws.com"

	c := colly.NewCollector(
		colly.AllowedDomains(o, "ec2co-ecsel-ldr54y7zto1q-1944992330.ap-northeast-1.elb.amazonaws.com"),
	)

	c.Limit(&colly.LimitRule{Delay: 1 * time.Second})

	// err := c.Post("http://ec2co-ecsel-ldr54y7zto1q-1944992330.ap-northeast-1.elb.amazonaws.com:8008/api-token-auth/", map[string]string{"username": "eee@gmail.com", "password": "pass1234"})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	c.OnHTML("a[href", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fmt.Println("link found ", e.Text, link)
		fmt.Println(e)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("res ", r.StatusCode)
	})

	fmt.Println("start")
	c.Visit("http://" + o)
	fmt.Println("end")
}
