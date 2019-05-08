package main

import (
	"context"
	"fmt"

	elastic "github.com/olivere/elastic"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	fmt.Println("conn es succ")

	var count = 10000
	for i := 2; i < count; i++ {
		tweet := Tweet{User: "olivere", Message: "Take Five"}
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id(fmt.Sprintf("%d", i)).
			BodyJson(tweet).
			Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
			return
		}
	}

	fmt.Println("insert succ")
}
