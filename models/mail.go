package models

// Mail model
type Mail struct {
	Sender   string
	Password string
	SmtpHost string
	SmtpPort string
}

func SetMailConfig(info map[string]string) (mail Mail) {
	mail.Sender = info["SENDER"]
	mail.Password = info["PASSWORD"]
	mail.SmtpHost = info["SMTP_HOST"]
	mail.SmtpPort = info["SMTP_PORT"]
	return mail
}
