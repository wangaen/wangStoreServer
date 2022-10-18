package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

func ItemSave() chan interface{} {
	//client, err := elastic.NewClient(elastic.SetSniff(false))
	//if err != nil {
	//	return nil, err
	//}
	out := make(chan interface{})

	go func() {
		itemCount := 0

		for {
			item := <-out
			log.Printf("保存第 %v 个 （%v） ", itemCount, item)
			//save(client, item)
			itemCount++
		}
	}()

	return out
}

func save(client *elastic.Client, item interface{}) error {
	//if item.Type == "" {
	//	return errors.New("item 类型为空")
	//}

	indexService := client.Index().Index("dating_profile").Type("").BodyJson(item)

	//if item.Id != "" {
	//	indexService.Id(item.Id)
	//}

	_, err := indexService.Do(context.Background())

	if err != nil {
		panic(err)
	}
	return nil
}
