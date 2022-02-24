package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
	"safakkizkin/config"
	"safakkizkin/migrations"
	"safakkizkin/routers"
)

var err error

func main() {
	handler()
}

func handler() {
	result := getEnviroment()
	dbName := result["DB_NAME"]
	dbPass := result["DB_PASS"]
	dbHost := result["DB_HOST"]
	dbPort := result["DB_PORT"]
	dbUser := result["DB_USER"]
	config.DB, err = gorm.Open("mysql", dbUser+":"+dbPass+"@("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	defer config.DB.Close()
	migrations.InitialMigration()

	r := routers.SetupRouters()
	r.Run(":3001")
}

// GetEnviroment to get env.
func getEnviroment() map[string]string {
	jsonFile, err := os.Open("environment.json")
	if err != nil {
		fmt.Println("File is not exist or can not read it. err:", err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)
	return result
}
