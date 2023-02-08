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

const redisAddr = "login-apollo.dev.cogoport.io:6379"

func StartClient(ID uint, Type string) {
	log.Print("Start of Ticket Client")
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: "f7d8279ad6ecaea58ccffd277a79b1cc4019da22713118805a9341d15a76c178",
	})

	tat := "00d:00h:05m"
	Duration := helpers.GetDuration(tat)

	if Type == "expiration" {
		task, err := tasks.ScheduleTicketExpirationTask(ID)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		info, err := client.Enqueue(task, asynq.ProcessIn(time.Duration(Duration)*time.Minute))
		if err != nil {
			log.Fatalf("could not enqueue expiration task: %v", err)
		}
		log.Print("Task done", info)

	} else if Type == "escalation" {
		task, err := tasks.ScheduleTicketEscalationTask(ID)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		info, err := client.Enqueue(task, asynq.ProcessIn(time.Duration(Duration)*time.Minute))
		if err != nil {
			log.Fatalf("could not enqueue escalation task: %v", err)
		}
		log.Print("Task done", info)
	} else if Type == "communication" {
		task, err := tasks.ScheduleTicketCommunicationTask(ID)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		info, err := client.Enqueue(task, asynq.ProcessIn(time.Duration(Duration)*time.Minute))
		if err != nil {
			log.Fatalf("could not enqueue expiration task: %v", err)
		}
		log.Print("Task done", info)
	}

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
