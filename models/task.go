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
	Name           string        `json:"Name"`
	UserId         uint          `json:"UserId"`
	UserMail       string        `json:"UserMail"`
	StartDate      time.Time     `json:"StartDate"`
	EndDate        time.Time     `json:"EndDate"`
	ReminderPeriod time.Duration `json:"ReminderPeriod"`
	LastRemindTime time.Time     `json:"LastRemindTime"`
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

func GetAllTasksByDate(t *[]Task) (err error) {
	if err = config.DB.Where("end_date > ?", time.Now()).Find(t).Error; err != nil {
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
		t1 := eachTask.StartDate.After(t.StartDate) && eachTask.EndDate.Before(t.StartDate)
		t2 := eachTask.StartDate.After(t.EndDate) && eachTask.EndDate.Before(t.EndDate)
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

func UpdateReminderDate(u *Task) (err error) {
	if err := config.DB.Model(&User{}).Where("ID = ?", u.Model.ID).Update("last_reminder_time", time.Now()).Error; err != nil {
		return err
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
