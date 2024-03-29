package app

import (
	"time"
	"github.com/LiamYabou/top100-ranking/variable"
	"github.com/getsentry/sentry-go"
)

func Finalize() {
	DBpool.Close()
	AMQPconn.Close()
	if variable.Env == "development" {
		file.Close()
	} else {
		sentry.Recover() // Capture the unhandled panic
  		sentry.Flush(5 * time.Second) // Set the timeout to the maximum duration the program can afford to wait.
	}
}
