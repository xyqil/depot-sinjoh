import (
  "github.com/millefalcon/go-subprocess/subprocess",
  "github.com/go-co-op/gocron/"
)
func main() {
  s := gocron.NewScheduler(time.UTC)
  // cron expressions supported
  s.Cron("*/30 * * * *").Do(RunDownloadTask()) // every minute
  // you can start running the scheduler in two different ways:
  // starts the scheduler asynchronously
  s.StartAsync()
  // starts the scheduler and blocks current execution path 
  s.StartBlocking()
  return nil
}
func RunDownloadTask() {
  proc := subprocess.Popen("ls")
  fmt.Println(proc.Stdout.Read())
  return nil
}
