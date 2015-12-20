package main

import (
  "fmt"
  "os"
  "os/signal"
  "os/exec"
  "log"
  "syscall"
)

// Actuall we don't really need this
func exitOnSignal() {
  sigs := make(chan os.Signal, 1);
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM);
  go func() {
    <- sigs;
    fmt.Println();
    os.Exit(1);
  }()
}

func main() {
  fmt.Println("Say something, press 'ctrl + c' or type 'exit' to exit");

  exitOnSignal();

  var command string;
  for command != "exit" {
    fmt.Printf("> ");
    fmt.Scanln(&command);
    cmd := exec.Command("say", command);
    err := cmd.Run();
    if err != nil {
      log.Fatal(err);
    }
    // TODO
  }

}
