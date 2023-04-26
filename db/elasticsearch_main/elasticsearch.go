package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"simple-video-server/config"
)

//func main() {
//	client, err := elasticsearch.NewClient(elasticsearch.Config{
//		Addresses: []string{"http://192.168.10.105:9200"},
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	info, err := client.Info()
//	if err != nil {
//		log.Fatalf("Error getting response: %s ", err)
//	}
//
//	defer info.Body.Close()
//
//	log.Println("info ", info)
//}

func main() {
	ctx := context.Background()
	//url := "http://192.168.10.105:9200"
	url := fmt.Sprintf("http://%s:%d", config.Elasticsearch.Host, config.Elasticsearch.Port)

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(url))
	if err != nil {
		panic(err)
	}

	_, _, err = client.Ping(url).Do(ctx)
	if err != nil {
		log.Println("连接es失败 ", err)
	}

	version, err := client.ElasticsearchVersion(url)
	if err != nil {
		log.Println("查询 es版本失败 ", err)
	}

	log.Println("es version ", version)
}
