package tools

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 官方事务文档：https://docs.mongodb.com/manual/core/transactions/

type Session struct {
	Client         *mongo.Client
	collection     *mongo.Collection
	uri            string
	Options        options.IndexOptions
	FindOneOptions options.FindOneOptions
	FindOptions    options.FindOptions
}

func New(uri string) *Session {
	session := &Session{
		uri: uri,
	}
	return session
}

// 初始化数据库
func (s *Session) InitMongoDB() error {
	var ClientOpts = options.Client().
		// 基本设置
		SetConnectTimeout(10 * time.Second).     // 连接超时
		SetHosts([]string{"10.100.0.31:27017"}). // 指定服务器地址
		SetMaxPoolSize(10).                      // 连接池连接数 - 最大
		SetMinPoolSize(1)                        // 连接池连接数 - 最小

	// 创建客户端FindOneOptions
	client, err := mongo.Connect(context.TODO(), ClientOpts)
	if err != nil {
		return err
	}
	s.Client = client
	return nil
}

// 创建索引（可以弃用了）
// 参考：https://stackoverflow.com/questions/56759074/how-do-i-create-a-text-index-in-mongodb-with-golang-and-the-mongo-go-driver
func (s *Session) AddIndex(dbName string, collectionName string, indexKeys interface{}) {

	// 手动实现，可以用
	// aaa := options.Index()
	// aaa.SetUnique(true)

	coll := s.Client.Database(dbName).Collection(collectionName)
	indexName, err := coll.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    indexKeys,
		Options: &s.Options,
		// Options: options.Index().SetUnique(true),	// 原始格式
	})
	if err != nil {
		log.Printf("创建索引失败，DBName：%v, collectionName：%v, indexKeys：%v", dbName, collectionName, indexKeys)
		log.Panic(err)
	}
	log.Printf("创建索引成功：%v", indexName)
}

// 创建单字段索引（优化使用）
// 参考：https://stackoverflow.com/questions/56759074/how-do-i-create-a-text-index-in-mongodb-with-golang-and-the-mongo-go-driver
func (s *Session) AddIndexSingle(dbName string, collectionName string, Key string, Desc int, SetUnique bool) {

	// 设置索引 Key
	indexKeys := map[string]interface{}{
		Key: Desc,
	}

	// 是否设置为唯一索引
	if SetUnique {
		options.Index().SetUnique(SetUnique)
	}

	// 设置索引名称
	s.Options.SetName(Key + "_" + fmt.Sprintln(Desc))

	// 开始设置索引
	coll := s.Client.Database(dbName).Collection(collectionName)
	indexName, err := coll.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    indexKeys,
		Options: &s.Options,
		// Options: options.Index().SetUnique(true),	// 原始格式
	})

	// mongo.ErrInvalidIndexValue
	// 如果有问题，直接报错
	if err != nil {

		if e, ok := err.(mongo.ServerError); ok {
			fmt.Println(e.Error())
			if e.HasErrorCode(85) {
				log.Warnf("索引已经存在%v\n", e)
				return
			}
		}

		log.Errorf("创建索引失败，DBName：%v, collectionName：%v, Key：%v, Desc：%v, SetUnique：%v", dbName, collectionName, Key, Desc, SetUnique)
		log.Panic(err)
	}

	// 否则的话输出创建索引的信息
	log.Printf("创建索引成功：%v", indexName)
}

// 插入一条数据
func (s *Session) InsertOne(dbName, collectionName string, doc interface{}, ctx ...context.Context) *mongo.InsertOneResult {
	coll := s.Client.Database(dbName).Collection(collectionName)

	result, err := coll.InsertOne(context.TODO(), doc)
	// result, err := coll.InsertOne(ctx[0], doc)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Println("主键冲突")
			return nil
		}
		log.Panicln(err)
	}
	return result
}

// 插入多条数据
func (s *Session) InsertMany(dbName, collectionName string, doc []interface{}) *mongo.InsertManyResult {
	coll := s.Client.Database(dbName).Collection(collectionName)

	result, err := coll.InsertMany(context.TODO(), doc)
	if err != nil {
		if !mongo.IsDuplicateKeyError(err) { // 如果不是主键冲突的错误，引发恐慌
			log.Panicln(err)
		}
	}

	// totalCount := len(doc)
	// count := len(result.InsertedIDs)
	// log.Printf("传入文档数量：%v, 插入文档数量：%v", totalCount, count)
	return result
}

// 查找一条数据
func (s *Session) FindOne(dbName, collectionName string, filter interface{}, ret interface{}, opts ...*options.FindOneOptions) (notFind bool, err error) {
	// func (s *Session) FindOne(dbName, collectionName string, filter interface{}, ret interface{}) (notFind bool, err error) {
	coll := s.Client.Database(dbName).Collection(collectionName)
	// options.FindOneOptions.SetSort()
	// opts := options.FindOne().SetSort(bson.D{{"number", -1}})
	err = coll.FindOne(context.TODO(), filter, opts...).Decode(ret)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			notFind = true
			err = nil
		}
	}
	return
}

// 查找多条数据
func (s *Session) Find(dbName, collectionName string, filter interface{}, opts ...*options.FindOptions) (results []bson.M, err error) {
	coll := s.Client.Database(dbName).Collection(collectionName)
	cursor, err := coll.Find(context.TODO(), filter, opts...)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
		}
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results, err
}
