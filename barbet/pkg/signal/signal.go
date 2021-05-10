package signal

import (
  "log"
  "os"
  "os/signal"
)

func WaitForTermination() {
  intChannel := make(chan os.Signal, 1)
  signal.Notify(intChannel, _sigINT, _sigTERM)
  <-intChannel
  log.Println("Shutting down...")
}
