package main

import (
	"fmt"
	"os"

	"github.com/juev/rss-parser/helpers"
	"github.com/juev/rss-parser/rss"
)

func main() {
	if len(os.Args) < 2 {
		helpers.Usage()
		os.Exit(0)
	}
	feedURL := os.Args[1]

	resp, err := helpers.ReadWithClient(feedURL)
	if err != nil {
		helpers.Exit("failed to read %s: %v", feedURL, err)
	}
	defer resp.Body.Close()

	channel, err := rss.ParseFeed(resp)
	if err != nil {
		helpers.Exit("cannot parse feed url: %v", err)
	}

	for _, item := range channel.Item {
		fmt.Println(item.Link)
	}
}
