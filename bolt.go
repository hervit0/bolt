package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "strings"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.EnableBashCompletion = true

  app.Name = "Bolt"
  app.Usage = "Bolt is wrapping various helpers"
  app.Action = func(c *cli.Context) error {
    fmt.Println("This is Bolt, try `bolt help` for more")
    return nil
  }

  app.Commands = []cli.Command{
    {
      Name:    "logs",
      Aliases: []string{"l"},
      Usage:   "See the current git logs of your folder",
      Action:  func(c *cli.Context) error {
        commandLine := strings.Split("git log --graph --abbrev-commit --decorate", " ")
        out, err := exec.Command(commandLine[0], commandLine[1:]...).Output()

        if err != nil {
          log.Fatal(err)
        }

        fmt.Printf("%s", out)
        return nil
      },
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
