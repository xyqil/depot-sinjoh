import (
    "github.com/flashmob/go-guerrilla";
    "log";
)
func main() {
  daemon := guerrilla.Daemon{};
  err := daemon.Start();
  if err == nil {
    log.Print("Server Started!");
  };
};
