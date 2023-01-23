package workers

import (
	"log"
	"strconv"
	"strings"
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

func GetDuration(ExpiryDuration string) int {
	duration := strings.Split(ExpiryDuration, ":")

	durationd := strings.Split(duration[0], "d")
	durationh := strings.Split(duration[1], "h")
	durationm := strings.Split(duration[2], "m")

	d, _ := strconv.Atoi(durationd[0])
	h, _ := strconv.Atoi(durationh[0])
	m, _ := strconv.Atoi(durationm[0])

	h += m / 60
	h += d * 24

	return h

}

// const redisAddr = "login-apollo.dev.cogoport.io:6379"

func StartTicketClient() {
	log.Print("Start of Ticket Client")
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: "f7d8279ad6ecaea58ccffd277a79b1cc4019da22713118805a9341d15a76c178",
	})

	task1, err := tasks.ScheduleTicketEscalationTask(18, 51, 1, 45, 51)

	tat := "00d:00h:05m"

	Duration := GetDuration(tat)

	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info1, err1 := client.Enqueue(task1, asynq.ProcessIn(time.Duration(Duration)*time.Minute))


	if err1 != nil {
		log.Fatalf("could not enqueue task: %v", err1)
	}

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
	mux.HandleFunc(tasks.TicketEscalation, tasks.HandleTicketEscalationTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
	log.Print("Handle Email task")

}
