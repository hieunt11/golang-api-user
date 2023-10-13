package connectdb

import (
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"
)

type ConfigDB struct {
	Database struct {
		Host            string `yaml: "host"`
		Port            int    `yaml: "port"`
		Username        string `yaml: "username"`
		Password        string `yaml: "password"`
		Databasename    string `yaml: "databasename"`
		Maxconnects     int    `yaml: "maxconnect"`
		Maxidleconnects int    `yaml: "maxidleconnect"`
	} `yaml:"database"`
}

func ConnectDB() (*sqlx.DB, error) {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")

	if err != nil {
		log.Fatalf("Error reading the YAML file: %v", err)
	}

	var config ConfigDB

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Databasename)

	// Connect DB
	db, err := sqlx.Connect("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	// Set the maximum number of open and idle connections in the pool
	db.SetMaxOpenConns(config.Database.Maxconnects)     // Maximum open connections
	db.SetMaxIdleConns(config.Database.Maxidleconnects) // Maximum idle connections

	fmt.Println("Successfully connected DB!")

	return db, nil
}
