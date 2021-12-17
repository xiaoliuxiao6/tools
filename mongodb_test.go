package tools_test

import (
	"fmt"
	"log"

	"testing"

	"github.com/xiaoliuxiao6/tools"
	"go.mongodb.org/mongo-driver/bson"
)

type Blog struct {
	Name string
	Age  int
}

func TestUsaged(t *testing.T) {

	dbName := "insertDB"
	collectionName := "haikus"

	// 建立连接
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

	// 插入结构体 - 单条数据
	bloger := Blog{
		Name: "Tom",
		Age:  18,
	}
	mongoClient.InsertOne(dbName, collectionName, bloger)

	// 插入结构体 - 多条数据
	bloger1 := Blog{
		Name: "Tom1",
		Age:  18,
	}
	bloger2 := Blog{
		Name: "Tom2",
		Age:  18,
	}
	blogers := make([]interface{}, 0)
	blogers = append(blogers, bloger1)
	blogers = append(blogers, bloger2)
	fmt.Println(blogers)
	mongoClient.InsertMany(dbName, collectionName, blogers)

	// 插入单个文档
	doc := bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}}
	mongoClient.InsertOne(dbName, collectionName, doc)

	// 插入多个额文档
	docs := []interface{}{
		bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}},
		bson.D{{"title", "Showcasing a Blossoming Binary"}, {"text", "Binary data, safely stored with GridFS. Bucket the data"}},
	}
	mongoClient.InsertMany(dbName, collectionName, docs)
}
