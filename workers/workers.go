package workers

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/tejas-cogo/go-cogoport/tasks"
)

const redisAddr = "login-apollo.dev.cogoport.io:637,"

func Workers() {
	log.Print("Start of Worker")
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr, Username: "apollo_app", Password: "f7d8279ad6ecaea58ccffd277a79b1cc4019da22713118805a9341d15a76c178"},
		asynq.Config{Concurrency: 10},
	)
	log.Print("End of New server")
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeWelcomeEmail, tasks.HandleWelcomeEmailTask)

	log.Print("Mux created")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
