package helpers

import (
	"safakkizkin/models"
	"time"
)

var tasks []models.Task

func NotifyByMail(task models.Task) {
	// TODO (@Safakkizkin): Do the mailing things in here.
}

func SetReminder() {
	err := models.GetAllTasks(&tasks)
	if err != nil {
		return
	}
}

func Check() {
	for ;; {
		// Sleep for 30 seconds every for loops ends.
		time.Sleep(30 * time.Second)

		for index, eachTask := range tasks {
			addedLastRemindTime := eachTask.LastRemindTime.Add(eachTask.ReminderPeriod * time.Minute)
			if time.Now().After(addedLastRemindTime) {
				go NotifyByMail(eachTask)
				tasks[index].LastRemindTime = time.Now()
				models.UpdateReminderDate(&eachTask)
			}
		}
	}
}