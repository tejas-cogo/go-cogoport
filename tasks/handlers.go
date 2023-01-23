package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

type EmailDeliveryPayload struct {
	user_id int
}

func HandleWelcomeEmailTask(c context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User: user_id=%d", p.user_id)
	return nil
}

// HandleReminderEmailTask for reminder email task.
func HandleReminderEmailTask(c context.Context, t *asynq.Task) error {
    // Get int with the user ID from the given task.
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
        return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
    }
    log.Printf("Sending Email to User: user_id=%d", p.user_id)
	return nil
}
