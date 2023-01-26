package tasks

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TicketEscalation = "ticket:escalated"
	TicketExpiration = "ticket:expiration"
)

type TicketExpirationPayload struct {
	TicketID       uint
	ReviewerUserID uint
	GroupID        uint
	GroupMemberID  uint
	GroupHeadID    uint
	Tat            time.Time
	ExpiryDate     time.Time
}

type TicketEscalationPayload struct {
	TicketID       uint
	ReviewerUserID uint
	GroupID        uint
	GroupMemberID  uint
	GroupHeadID    uint
	Tat            time.Time
	ExpiryDate     time.Time
}

func ScheduleTicketEscalationTask(TicketID uint) (*asynq.Task, error) {
	payload, err := json.Marshal(TicketEscalationPayload{TicketID: TicketID})

	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TicketEscalation, payload), nil
}

func ScheduleTicketExpirationTask(TicketID uint, ReviewerUserID uint, GroupID uint, GroupMemberID uint, GroupHeadID uint) (*asynq.Task, error) {
	payload, err := json.Marshal(TicketExpirationPayload{TicketID: TicketID, ReviewerUserID: ReviewerUserID, GroupID: GroupID, GroupMemberID: GroupMemberID, GroupHeadID: GroupHeadID})

	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TicketExpiration, payload), nil
}

