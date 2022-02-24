package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"safakkizkin/config"
	"time"
)

// Task model
type Task struct {
	gorm.Model
	Name           string    `json:"Name"`
	UserId         uint      `json:"UserId"`
	UserMail       string    `json:"UserMail"`
	StartDate      time.Time `json:"StartDate"`
	EndDate        time.Time `json:"EndDate"`
	ReminderPeriod uint16    `json:"ReminderPeriod"`
}

func GetAllTasksByUser(t *[]Task, u *User) (err error) {
	if err = config.DB.Where("user_id = ?", u.Model.ID).Find(t).Error; err != nil {
		return err
	}

	return nil
}

func GetAllTasks(t *[]Task) (err error) {
	if err = config.DB.Find(t).Error; err != nil {
		return err
	}

	return nil
}

func checkTaskOverlap(u *User, t *Task) (err error) {
	var tasks []Task
	err = GetAllTasksByUser(&tasks, u)
	if err != nil {
		return err
	}

	isThereOverlap := false
	for _, eachTask := range tasks {
		t1 := eachTask.StartDate.Before(t.EndDate)
		t2 := eachTask.EndDate.After(t.EndDate)
		if t1 || t2 {
			isThereOverlap = true
			break
		}
	}

	if isThereOverlap {
		return errors.New("task: there is already a task assigned to this user")
	}

	return nil
}

func DeleteTask(u *Task, id string) (err error) {
	if err := config.DB.Where("ID = ?", id).Delete(u).Error; err != nil {
		return err
	}

	return nil
}

func GetTask(t *Task, id string) (err error) {
	if err := config.DB.Where("ID = ?", id).First(t).Error; err != nil {
		return err
	}

	return nil
}

func AddNewTask(t *Task) (err error) {
	var user User
	user.Mail = t.UserMail
	err = CheckIfUserPresent(&user)
	if err != nil {
		return
	}

	err = checkTaskOverlap(&user, t)
	if err != nil {
		return err
	}

	t.UserId = user.Model.ID
	if err = config.DB.Create(t).Error; err != nil {
		return err
	}

	return nil
}
