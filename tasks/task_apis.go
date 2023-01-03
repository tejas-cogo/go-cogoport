package tasks

import (
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

const (
	TypeWelcomeEmail  = "email:welcome"
	TypeReminderEmail = "email:reminder"
)

type EmailTaskPayloads struct {
	user_id int
}

type EmailPayload struct {
	user_id int
	sent_id string
}

func NewWelcomeEmailTask(id int) *asynq.Task {
	payload, err := json.Marshal(EmailTaskPayloads{user_id: id})

	if err != nil {
		log.Fatal(err)
	}
	return asynq.NewTask(TypeWelcomeEmail, payload)
}
