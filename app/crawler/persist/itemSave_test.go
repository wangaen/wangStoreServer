package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
	"wangStoreServer/app/crawler/engine"
	models "wangStoreServer/app/crawler/models/zhenaiwang"
)

var example = engine.Item{
	Url:  "https://album.zhenai.com/u/105824871",
	Type: "zhenaiwang",
	Id:   "105824871",
	Payload: models.User{
		Name:        "张三",
		Sex:         "男士",
		Age:         "33岁",
		Height:      "172cm",
		Weight:      "69kg",
		Salary:      "1.2-2万",
		Status:      "未婚",
		XingZuo:     "魔羯座(12.22-01.19)",
		XueLi:       "大学本科",
		Work:        "咨询/顾问",
		WorkAddress: "重庆南岸区",
		Signature:   "茫茫人海中，找一个知心，恩爱的她，希望这个她就是你，让我们一起搭建一个幸福美满的家！",
		GirlCondition: models.GirlCondition{
			Age:         "25-35岁",
			Height:      "150-164cm",
			Salary:      "3千以上",
			XueLi:       "大专",
			WorkAddress: "重庆",
		},
	},
}

func TestSave(t *testing.T) {
	err := save(example)
	if err != nil {
		t.Fatalf("save(example) err : %v\n", err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	getResult, err := client.Get().Index("dating_profile").Type(example.Type).Id(example.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	var data engine.Item
	err = json.Unmarshal(*getResult.Source, &data)
	if err != nil {
		panic(err)
	}

	if example != data {
		t.Errorf("两者不相同: \n example:%#v \n data:%#v\n", example, data)
	}
}
