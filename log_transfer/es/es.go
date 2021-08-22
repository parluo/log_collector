package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

type MsgData struct {
	Time time.Time
	Str  string
}

func Init(addr string) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	sniffOpt := elastic.SetSniff(false)
	client, err = elastic.NewClient(elastic.SetURL(addr), sniffOpt)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es...")
	return nil
}

func SendToEs(idx string, data []byte) {
	tmpMsg := MsgData{
		Time: time.Now(),
		Str:  string(data),
	}
	_, err := client.Index().Index(idx).BodyJson(tmpMsg).Do(context.Background())
	if err != nil {
		fmt.Println("sent msg to es failed...")
		return
	}

}
