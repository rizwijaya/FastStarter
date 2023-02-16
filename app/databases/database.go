package database

import (
	"FastStarter/app/config"
	"context"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database map[string]interface{}

func MySql(config config.LoadConfig) *gorm.DB {
	dsn := config.Database.Username + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.Name + "?charset=utf8&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}
	return Db
}

func Postgres(config config.LoadConfig) *gorm.DB {
	dsn := "host=" + config.Database.Host + " user=" + config.Database.Username + " password=" + config.Database.Password + " dbname=" + config.Database.Name + " port=" + config.Database.Port + " sslmode=disable TimeZone=Asia/Jakarta"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}
	return Db
}

func MongoDb(config config.LoadConfig) *mongo.Database {
	var clientOptions *options.ClientOptions
	ctx := context.TODO()
	if config.Database.Username == "" && config.Database.Password == "" {
		clientOptions = options.Client().ApplyURI("mongodb://" + config.Database.Host + ":" + config.Database.Port)
	} else {
		clientOptions = options.Client().ApplyURI("mongodb://" + config.Database.Username + ":" + config.Database.Password + "@" + config.Database.Host + ":" + config.Database.Port)
	}
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	db := client.Database(config.Database.Name)
	return db
}

func NewDatabase() (map[string]interface{}, error) {
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	database = make(map[string]interface{})

	if config.Database.Driver == "mysql" {
		database["type"] = "mysql"
		database["connection"] = MySql(config)
	} else if config.Database.Driver == "postgres" {
		database["type"] = "postgres"
		database["connection"] = Postgres(config)
	} else if config.Database.Driver == "mongoDb" {
		database["type"] = "mongoDb"
		database["connection"] = MongoDb(config)
	} else {
		database["type"] = "none"
		database["connection"] = nil
		log.Fatalln("Database driver not found")
	}
	return database, nil
}
