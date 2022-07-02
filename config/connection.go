package config

import (
	"Test-Rizky/domain"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//get data from yml

type Config struct {
	Database struct {
		DBName string `yaml:"db"`
		Redis  struct {
			RedisAddress  string `yaml:"url"`
			RedisPassword string `yaml:"password"`
		} `yaml:"redis"`
	} `yaml:"database"`
}

//check error
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//connection db

func Connection() *gorm.DB {
	fmt.Println("yaml config")

	var config Config

	file, err := os.Open("local.yaml")

	check(err)

	defer file.Close()

	content, err := ioutil.ReadAll(file)

	check(err)

	err = yaml.Unmarshal(content, &config)

	check(err)

	dsn := config.Database.DBName
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("connection failed", err)
	} else {
		log.Println("connection db success")
	}

	err = db.AutoMigrate(domain.Order{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(domain.Customer{})
	if err != nil {
		return nil
	}
	return db
}
