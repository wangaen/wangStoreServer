package persist

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"log"
	"wangStoreServer/app/crawler/engine"
)

func ItemSave() chan engine.Item {
	out := make(chan engine.Item)

	go func() {
		itemCount := 0

		for {
			item := <-out
			err := save(item)
			if err != nil {
				log.Printf("保存item出现异常：保存 （%v）,err: （%v）\n", item, err.Error())
			} else {
				log.Printf("保存第 %v 个 （%v）\n", itemCount, item)
			}
			itemCount++
		}
	}()

	return out
}

func save(item engine.Item) error {

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("elastic type is null")
	}

	indexService := client.Index().Index("dating_profile").Type(item.Type)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
