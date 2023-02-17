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

type Database interface {
	Type() string
	Mysql() *gorm.DB
	Postgres() *gorm.DB
	Mongodb() *mongo.Database
}

type database struct {
}

func NewDatabase() *database {
	return &database{}
}

func (d *database) Type() string {
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return config.Database.Driver
}

func (d *database) Mysql() *gorm.DB {
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	dsn := config.Database.Username + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.Name + "?charset=utf8&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}
	return Db
}

func (d *database) Postgres() *gorm.DB {
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	dsn := "host=" + config.Database.Host + " user=" + config.Database.Username + " password=" + config.Database.Password + " dbname=" + config.Database.Name + " port=" + config.Database.Port + " sslmode=disable TimeZone=Asia/Jakarta"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}
	return Db
}

func (d *database) Mongodb() *mongo.Database {
	var clientOptions *options.ClientOptions
	ctx := context.TODO()
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
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
