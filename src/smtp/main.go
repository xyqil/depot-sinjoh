import (
    "github.com/flashmob/go-guerrilla";
    "log";
)
func main() {
  d := guerrilla.Daemon{};
  err := d.Start();
  if err == nil {
    log.Print("Server Started!");
  };
};
