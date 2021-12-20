package tools

import (
	"context"
	"log"
	"time"
    "errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Session struct {
	client     *mongo.Client
	collection *mongo.Collection
	uri        string
}

func New(uri string) *Session {
	session := &Session{
		uri: uri,
	}
	return session
}

func (s *Session) InitMongoDB() error {
	var ClientOpts = options.Client().
		// 基本设置
		SetConnectTimeout(10 * time.Second).     // 连接超时
		SetHosts([]string{"10.100.0.31:27017"}). // 指定服务器地址
		SetMaxPoolSize(10).                      // 连接池连接数 - 最大
		SetMinPoolSize(1)                        // 连接池连接数 - 最小

	// 创建客户端
	client, err := mongo.Connect(context.TODO(), ClientOpts)
	if err != nil {
		return err
	}
	s.client = client
	return nil
}

// 判断是否为主键重复
func IsDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}

// 插入一条数据
func (s *Session) InsertOne(dbName, collectionName string, doc interface{}) {
	coll := s.client.Database(dbName).Collection(collectionName)

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		if IsDup(err) {
			log.Println("主键重复")
		}
		log.Panicln(err)
	}

	log.Printf("插入文档的 ID：%v\n", result.InsertedID)
}

// 插入多条数据
func (s *Session) InsertMany(dbName, collectionName string, doc []interface{}) {
	coll := s.client.Database(dbName).Collection(collectionName)

	result, err := coll.InsertMany(context.TODO(), doc)
	if err != nil {
		log.Panicln(err)
	}

	count := len(result.InsertedIDs)
	log.Printf("插入文档数量：%v", count)
}

// // 查找数据
// func (s *Session) FindOne(dbName, collectionName string, doc []interface{}) {
// 	coll := s.client.Database(dbName).Collection(collectionName)
// 	result := coll.FindOne(context.TODO(), doc)
// 	log.Printf("插入文档数量：%v", result)
// }

