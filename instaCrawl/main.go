package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var requestID string
var requestIds [][]byte
var queryIdPattern = regexp.MustCompile(`queryId:".{32}"`)

type pageInfo struct {
	EndCursor string `json:"end_cursor"`
	NextPage  bool   `json:"has_next_page"`
}

func main() {
	var actualUserId string
	url := "https://www.instagram.com/tk_streetphotography"

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// c.Post()

	c.Limit(&colly.LimitRule{Delay: 1 * time.Second})

	c.OnHTML("body > script:first-of-type", func(e *colly.HTMLElement) {
		jsonData := e.Text[strings.Index(e.Text, "{") : len(e.Text)-1]

		// parse JSON
		data := struct {
			EntryData struct {
				ProfilePage []struct {
					User struct {
						Id    string `json:"id"`
						Media struct {
							Nodes []struct {
								ImageURL     string `json:"display_src"`
								ThumbnailURL string `json:"thumbnail_src"`
								IsVideo      bool   `json:"is_video"`
								Date         int    `json:"date"`
								Dimensions   struct {
									Width  int `json:"width"`
									Height int `json:"height"`
								}
							}
							PageInfo pageInfo `json:"page_info"`
						} `json:"media"`
					} `json:"user"`
				} `json:"ProfilePage"`
			} `json:"entry_data"`
		}{}
		err := json.Unmarshal([]byte(jsonData), &data)
		if err != nil {
			log.Fatal(err)
		}

		// enumerate images
		page := data.EntryData.ProfilePage[0]
		actualUserId = page.User.Id
		for _, obj := range page.User.Media.Nodes {
			// skip videos
			if obj.IsVideo {
				continue
			}
			c.Visit(obj.ImageURL)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("Referer", url)
		fmt.Println("visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("res ", r.StatusCode)
	})

	c.Visit(url)
	// c.Visit("https://github.com/gocolly/colly")
}
