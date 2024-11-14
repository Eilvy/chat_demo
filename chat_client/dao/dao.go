package dao

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go_code/chat_demo/chat_client/model"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func Init() {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	Dbname := "chat"
	timeout := "10s"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		log.Panicf("Connect DB error: %v  \n", err)
		return
	}
	DB = db
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("init user table error :", err)
		return
	}
	err = DB.AutoMigrate(&model.Message{})
	if err != nil {
		log.Println("init message table error : ", err)
		return
	}
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:        "redis-14520.c299.asia-northeast1-1.gce.cloud.redislabs.com:14520",
		Password:    "rPYdtUeiD5CeJSqcGZdoyHDd6Ou2uApa",
		DB:          0,
		DialTimeout: time.Second * 5,
	})
	RDB = client
}
