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
    fmt.Println("This is Bolt, try `bolt help` for more!")
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
    {
      Name:    "install",
      Aliases: []string{"i"},
      Usage:   "Install various useful tools",
      Subcommands: []cli.Command{
        {
          Name:  "docker-sync",
          Aliases: []string{"ds"},
          Usage: "Install docker-sync\t -\t a Ruby version manager is required",
          Action: func(c *cli.Context) error {
            cmd := "curl -s https://raw.githubusercontent.com/hervit0/bolt/master/scripts/install_docker_sync | sh"
            out, err := exec.Command("bash", "-c", cmd).Output()

            if err != nil {
              log.Fatal(err)
            }

            fmt.Printf("%s", out)
            return nil
          },
        },
        {
          Name:  "dex",
          Aliases: []string{"dx"},
          Usage: "Install dex\t -\t the gem `docker-sync` is required",
          Action: func(c *cli.Context) error {
            cmd := "curl -s https://raw.githubusercontent.com/hervit0/bolt/master/scripts/install_dex | sh"
            out, err := exec.Command("bash", "-c", cmd).Output()

            if err != nil {
              log.Fatal(err)
            }

            fmt.Printf("%s", out)
            return nil
          },
        },
      },
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
