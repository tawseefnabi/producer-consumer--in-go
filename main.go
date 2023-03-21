//////////////////////////////////////////////////////////////////////
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return tweets
		}
		tweets = append(tweets, tweet)
	}
}
func consumer(tweets []*Tweet) {
	for _, t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang", "\t\t", t.Text)
		} else {
			fmt.Println(t.Username, "\tdoes not tweets about golang")
		}
	}
}
func main() {
	start := time.Now()
	fmt.Println("start", start)
	stream := GetMockStream()
	// fmt.Println("stream", stream)
	// producer
	tweets := producer(stream)
	fmt.Println("tweets", tweets)
	// consumer
	consumer(tweets)
}
