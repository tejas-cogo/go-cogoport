package workers

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
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

// const redisAddr = "login-apollo.dev.cogoport.io:6379"

func StartTicketClient() {
	log.Print("Start of Ticket Client")
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: "f7d8279ad6ecaea58ccffd277a79b1cc4019da22713118805a9341d15a76c178",
	})

	task2, err := tasks.ScheduleTicketExpirationTask(18)
	task1, err := tasks.ScheduleTicketEscalationTask(18)

	tat := "00d:00h:05m"

	Duration := helpers.GetDuration(tat)

	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info2, err2 := client.Enqueue(task2, asynq.ProcessIn(time.Duration(Duration)*time.Minute))
	info1, err1 := client.Enqueue(task1, asynq.ProcessIn(time.Duration(Duration)*time.Minute))
	
	if err1 != nil {
		log.Fatalf("could not enqueue escalation task: %v", err1)
	}
	if err2 != nil {
		log.Fatalf("could not enqueue expiration task: %v", err2)
	}

	log.Print("Task done", info2)
	log.Print("Task done", info1)
	log.Print("End of New server")
}

func StartTicketHandler() {

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
	mux.HandleFunc(tasks.TicketExpiration, tasks.HandleTicketExpirationTask)
	mux.HandleFunc(tasks.TicketEscalation, tasks.HandleTicketEscalationTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
	log.Print("Handle Email task")

}
