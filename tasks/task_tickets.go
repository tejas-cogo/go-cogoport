package tasks

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

const (
	ScheduleTicket = "ticket:schedule"
	HandleTicket   = "ticket:escalated"
)

type TicketTaskPayloads struct {
	TicketID       uint
	ReviewerUserID uint
	GroupID        uint
	GroupMemberID  uint
	GroupHeadID    uint
	Tat            time.Time
	ExpiryDate     time.Time
}

type TicketPayload struct {
	TicketID       uint
	ReviewerUserID uint
	GroupID        uint
	GroupMemberID  uint
	GroupHeadID    uint
	Tat            time.Time
	ExpiryDate     time.Time
}

func ScheduleTicketTask(TicketID uint, ReviewerUserID uint, GroupID uint, GroupMemberID uint, GroupHeadID uint) (*asynq.Task, error) {
	payload, err := json.Marshal(TicketTaskPayloads{TicketID: TicketID, ReviewerUserID: ReviewerUserID, GroupID: GroupID, GroupMemberID: GroupMemberID, GroupHeadID: GroupHeadID})

	if err != nil {
		return nil, err
	}
	return asynq.NewTask(ScheduleTicket, payload), nil
}
