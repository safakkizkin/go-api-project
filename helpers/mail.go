package helpers

import (
	"fmt"
	"net/smtp"
	"safakkizkin/models"
)

func SendMail(auth smtp.Auth, mail *models.Mail, task models.Task) (err error) {
	messageString := fmt.Sprintf("Task %s will end at %s.\nPlease make sure you do the task on time.", task.Name, task.EndDate.String())
	message := []byte(messageString)
	err = smtp.SendMail(mail.SmtpHost+":"+mail.SmtpPort, auth, mail.Sender, []string{task.UserMail}, message)

	return err
}
