package services
import "github.com/robfig/cron/v3"

//https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc

type Scheduler struct {
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.start()

	// Inspect the cron job entries' next and previous run times.
	inspect(c.Entries())

	c.stop()
}
