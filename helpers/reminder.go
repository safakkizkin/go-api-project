package helpers

import (
	"fmt"
	"net/smtp"
	"safakkizkin/models"
	"time"
)

var tasks []models.Task
var needToRemove []int
var auth smtp.Auth
var mail models.Mail

func NotifyByMail(task models.Task, index int) {
	err := SendMail(auth, &mail, task)
	fmt.Println(err.Error())
	if err == nil {
		tasks[index].LastRemindTime = time.Now()
		models.UpdateReminderDate(&tasks[index])
	}
}

func SetMail(m models.Mail) (isMailSet bool) {
	if len(m.Sender) > 0 &&
		len(m.Password) > 0 &&
		len(m.SmtpHost) > 0 &&
		len(m.SmtpPort) > 0 {
		mail = m
		auth = smtp.PlainAuth("", m.Sender, m.Password, m.SmtpHost)
		return true
	}

	return false
}

func SetReminder() {
	err := models.GetAllTasks(&tasks)
	if err != nil {
		return
	}
}

func remove(slice []models.Task, s int) []models.Task {
	return append(slice[:s], slice[s+1:]...)
}

func Check() {
	for {
		for index, eachTask := range tasks {
			if time.Now().After(eachTask.EndDate) {
				needToRemove = append(needToRemove, index)
				continue
			}

			if time.Now().Before(eachTask.StartDate) {
				continue
			}

			addedLastRemindTime := eachTask.LastRemindTime.Add(eachTask.ReminderPeriod * time.Minute)
			if time.Now().After(addedLastRemindTime) {
				go NotifyByMail(eachTask, index)
			}
		}

		for i := len(needToRemove) - 1; i >= 0; i-- {
			tasks = remove(tasks,
				needToRemove[i])
		}

		needToRemove = []int{}
		time.Sleep(10 * time.Second)
	}
}
