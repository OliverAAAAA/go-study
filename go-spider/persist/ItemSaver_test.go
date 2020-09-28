package persist

import (
	"context"
	elastic "github.com/olivere/elastic/v7"
	"testing"
)

type item struct {
	name string
}

type Tweet struct {
	User string
	Message string
	Retweets int
}

func TestSaver(t *testing.T) {
	//item := item{name: "testSave"}
	//id, err := save(item)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(id)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	//if err != nil {
	//
	//}
	//do, err := client.Get().Index("spider_user").Type("zhenai").Id(id).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//t.Logf("%s\n",do.Source)

	// Index a second tweet (by string)
	tweet2 := Tweet{User: "Cat", Message: "Hello Five", Retweets: 0}
	_, err = client.Index().
		Index("spider_user").
//		Id("2").
		BodyJson(tweet2).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

}
