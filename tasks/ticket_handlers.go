package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/tejas-cogo/go-cogoport/models"
	worker "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/workers"
)

type TicketWorkerService struct {
	p models.TicketEscalatedPayload
}

// HandleReminderEmailTask for reminder email task.
func HandleTicketEscalationTask(c context.Context, t *asynq.Task) error {
	// Get int with the user ID from the given task.

	var p models.TicketEscalatedPayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	worker.TicketEscalation(p)

	log.Printf("Ticket Escalated=%d", p.TicketID)
	return nil
}
