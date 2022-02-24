package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
	"safakkizkin/config"
	"safakkizkin/helpers"
	"safakkizkin/migrations"
	"safakkizkin/models"
	"safakkizkin/routers"
)

var err error

func main() {
	handler()
}

func handler() {
	result := getEnvironment("environment.json")
	if result == nil {
		return
	}

	dbName := result["DB_NAME"]
	dbPass := result["DB_PASS"]
	dbHost := result["DB_HOST"]
	dbPort := result["DB_PORT"]
	dbUser := result["DB_USER"]
	config.DB, err = gorm.Open("mysql", dbUser+":"+dbPass+"@("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	mailConfig := getEnvironment("mailConfig.json")
	if mailConfig != nil {
		mail := models.SetMailConfig(mailConfig)
		isMailSet := helpers.SetMail(mail)
		defer config.DB.Close()
		migrations.InitialMigration()
		if isMailSet {
			helpers.SetReminder()
			go helpers.Check()
		}
	}

	r := routers.SetupRouters()
	err := r.Run(":3001")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// GetEnvironment to get env.
func getEnvironment(configFileName string) map[string]string {
	jsonFile, err := os.Open(configFileName)
	if err != nil {
		fmt.Println("File is not exist or can not read it. err:", err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]string
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}
