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

// 基本使用示例
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

	// 插入多个文档
	docs := []interface{}{
		bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}},
		bson.D{{"title", "Showcasing a Blossoming Binary"}, {"text", "Binary data, safely stored with GridFS. Bucket the data"}},
	}
	mongoClient.InsertMany(dbName, collectionName, docs)
}

// 插入单条数据
func TestInsertOne(t *testing.T) {

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
	result := mongoClient.InsertOne(dbName, collectionName, bloger)
	// log.Printf("插入文档的 ID：%v\n", result.InsertedID)
	// fmt.Println(len(result.InsertedID))
	if result == nil {
		fmt.Println("插入数量为空")
	}
}

// 插入多条数据
func TestInsertMany(t *testing.T) {

	dbName := "insertDB"
	collectionName := "haikus"

	// 建立连接
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

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
	// fmt.Println(blogers)
	mongoClient.InsertMany(dbName, collectionName, blogers)
}

// 创建索引
func TestAddIndex(t *testing.T) {
	// 建立连接
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

	// 设置是否为唯一索引
	mongoClient.Options.SetUnique(true)

	// // 单字段索引（方式1）
	// aaa := map[string]interface{}{
	// 	"myfieldname_type1": 1,
	// }
	// mongoClient.AddIndex("mydb", "mycollection111", aaa) // to descending set it to -1

	// // 单字段索引（方式2）
	// mongoClient.AddIndex("mydb", "mycollection111", bson.M{"myfieldname_type2": 1}) // to descending set it to -1

	// // 符合索引
	// mongoClient.AddIndex("mydb", "mycollection222", bson.D{{"myFirstField", 1}, {"mySecondField", -1}}) // to descending set it to -1

	// // 文本索引
	// mongoClient.AddIndex("mydb", "mycollection333", bson.D{{"myFirstTextField", "text"}, {"mySecondTextField", "text"}})

	// 插入多个文档
	// 符合索引
	mongoClient.AddIndex("mydb", "mycollection333", bson.D{{"myFirstField", 1}, {"mySecondField", -1}}) // to descending set it to -1

	// mongoClient.set
	docs := []interface{}{
		// 	bson.D{{"myFirstField", "aaa"}, {"mySecondField", "aaa"}, {"_id", "111"}},
		// 	bson.D{{"myFirstField", "bbb"}, {"mySecondField", "bbb"}, {"_id", "111"}},

		bson.D{{"myFirstField", "aaa"}, {"mySecondField", "aaa"}},
		bson.D{{"myFirstField", "aaa"}, {"mySecondField", "aaa"}},
	}
	mongoClient.InsertMany("mydb", "mycollection333", docs)
}

// 测试事务
func TestTransaction(t *testing.T) {

	// 建立连接
	dbName := "insertDB"
	collectionName := "haikus"
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

	type Stu struct {
		Name string `bson:"_id"`
		Age  int
	}

	// 插入结构体 - 单条数据
	bloger := Stu{
		Name: "Tom",
		Age:  20,
	}
	result := mongoClient.InsertOne(dbName, collectionName, bloger)
	if result == nil {
		fmt.Println("插入数量为空")
	} else {
		fmt.Println(result.InsertedID)
	}
}
