package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/iamr-kumar/go-service/internal/databases"
)

func startScrapping(db *databases.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Starting scrapping with concurrency %d and time between request %s", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feedsToFetch, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		
		if err != nil {
			log.Printf("Could not fetch feeds to fetch: %s", err)
			continue
		}

		waitGroup := &sync.WaitGroup{}

		for _, feed := range feedsToFetch {
			waitGroup.Add(1)
			go scrapeFeed(waitGroup, db, feed)
		}
		waitGroup.Wait()
	}
}

func scrapeFeed(waitGroup *sync.WaitGroup, db *databases.Queries, feed databases.Feed) {
	defer waitGroup.Done()

	log.Printf("Scrapping feed %s", feed.Url)
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Could not mark feed %d as fetched: %s", feed.ID, err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Printf("Could not fetch feed %s: %s", feed.Url, err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		log.Printf("Item: %s", item.Title)
	}

	log.Printf("Fetched feed %d", len(rssFeed.Channel.Item))

}