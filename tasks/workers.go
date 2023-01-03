package tasks

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/tejas-cogo/go-cogoport/tasks"
)

const redisAddr = "login-apollo.dev.cogoport.io:6379,"

func workers() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr, Password: "f7d8279ad6ecaea58ccffd277a79b1cc4019da22713118805a9341d15a76c178"},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDelivery, tasks.HandleEmailDeliveryTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
