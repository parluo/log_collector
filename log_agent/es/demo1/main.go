package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Person struct {
	OrderId  int    `json:"order_id"`
	OrderNum string `json:"order_num"`
	Addr     string `json:"addr"`
}

func main() {
	ctx := context.Background()
	sniffOpt := elastic.SetSniff(false) //panic: no active connection found: no Elasticsearch node available
	client, err := elastic.NewClient(elastic.SetURL("http://0.0.0.0:9200"), sniffOpt)
	if err != nil {
		// Handle error
		fmt.Printf("1:\n")
		panic(err)
	}
	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	fmt.Println("connect to es success")
	p1 := Person{1, "SPX000012", "beijing3"}
	put1, err := client.Index().
		Index("order_id").
		BodyJson(p1). // go中的对象转化为json
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Printf("2:\n")
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
