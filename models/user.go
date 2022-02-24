package models

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"safakkizkin/config"
)

// User model
type User struct {
	gorm.Model
	Mail      string `json:"Mail"`
	Firstname string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

func DeleteUser(u *User, mail string) (err error) {
	if err := config.DB.Where("Mail = ?", mail).Delete(u).Error; err != nil {
		return err
	}

	return nil
}

func GetUser(u *User) (err error) {
	if err := config.DB.Where("Mail = ?", u.Mail).First(u).Error; err != nil {
		return err
	}

	return nil
}

func GetAllUsers(u *[]User) (err error) {
	if err = config.DB.Find(u).Error; err != nil {
		return err
	}

	return nil
}

func AddNewUser(u *User) (err error) {
	if err = config.DB.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func CheckIfUserPresent(user *User) (err error) {
	if err := config.DB.Where("Mail = ?", user.Mail).First(user).Error; err != nil {
		return err
	}

	if user.Model.ID > 0 {
		return nil
	}

	return errors.New("user: not found")
}
