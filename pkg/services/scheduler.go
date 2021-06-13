package services
/*
Scheduler is using for scheduling scanning component within K8S ecosystem.
It takes cron expression to schedule tasks. 
*/
import (
	"fmt"

	"github.com/robfig/cron/v3"
)

//https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc

type Scheduler struct {
	cronExpression string
}

func StartScheduler(cronExpression string) {
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.Start()

	// Inspect the cron job entries' next and previous run times.
	//inspect(c.Entries())

}

func StopScheduler(c *cron.Cron) {
	c.Stop()
}
