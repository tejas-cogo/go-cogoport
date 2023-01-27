package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TicketEscalation = "ticket:escalated"
	TicketExpiration = "ticket:expired"
)

type TicketPayload struct {
	TicketID uint
}

func ScheduleTicketEscalationTask(TicketID uint) (*asynq.Task, error) {
	payload, err := json.Marshal(TicketPayload{TicketID: TicketID})

	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TicketEscalation, payload), nil
}

func ScheduleTicketExpirationTask(TicketID uint) (*asynq.Task, error) {
	payload, err := json.Marshal(TicketPayload{TicketID: TicketID})

	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TicketExpiration, payload), nil
}
