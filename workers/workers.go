package workers

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/tejas-cogo/go-cogoport/tasks"
)

const redisAddr = "login-apollo.dev.cogoport.io:6379"

func StartClient() {
	log.Print("Start of Client")
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: "f7d8279ad6ecaea58ccffd277a79b1cc4019da22713118805a9341d15a76c178",
	})

	task, err := tasks.NewWelcomeEmailTask(42)

	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task, asynq.ProcessIn(1*time.Minute))

	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
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
	mux.HandleFunc(tasks.TypeWelcomeEmail, tasks.HandleWelcomeEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
	log.Print("Handle Email task")

}
