package workers

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/tejas-cogo/go-cogoport/tasks"
)

const redisAddr = "login-apollo.dev.cogoport.io:6379"

func Workers() {
	log.Print("Start of Worker")
	client := asynq.NewClient(asynq.RedisClientOpt{
					Addr: redisAddr,
					Password: "f7d8279ad6ecaea58ccffd277a79b1cc4019da22713118805a9341d15a76c178",

				})
	defer client.Close()

	task, err := tasks.NewWelcomeEmailTask(42)

	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task)

	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}

	log.Print("Task done", info)
	log.Print("End of New server")
}
