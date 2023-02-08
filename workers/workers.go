package workers

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/tejas-cogo/go-cogoport/tasks"
)

type TicketData struct {
	TicketID       uint
	ReviewerUserID uint
	GroupID        uint
	GroupMemberID  uint
	GroupHeadID    uint
	Tat            time.Time
	ExpiryDate     time.Time
}

const redisAddr = "login-apollo.dev.cogoport.io:6379"

func StartClient(duration time.Duration, Task *asynq.Task) {
	log.Print("Start of Ticket Client")
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: "f7d8279ad6ecaea58ccffd277a79b1cc4019da22713118805a9341d15a76c178",
	})

	// Duration := helpers.GetDuration(Tat)

	info, err := client.Enqueue(Task, asynq.ProcessIn(duration))

	if err != nil {
		log.Fatalf("could not enqueue expiration task: %v", err)
	}
	log.Print("Task done", info)

	log.Print("End of New server")
}

func StartHandler() {

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr,
			Password: "f7d8279ad6ecaea58ccffd277a79b1cc4019da22713118805a9341d15a76c178",
		},
		asynq.Config{
			Concurrency: 10,
		},
	)
	log.Print("Starting Server")
	mux := asynq.NewServeMux()
	// mux.HandleFunc(tasks.TicketCommunication, tasks.HandleTicketCommunicationTask)
	mux.HandleFunc(tasks.TicketExpiration, tasks.HandleTicketExpirationTask)
	mux.HandleFunc(tasks.TicketEscalation, tasks.HandleTicketEscalationTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
	log.Print("Handle Email task")

}
