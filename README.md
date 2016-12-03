# cliconfig
Urfawe Cli flags setup from config struct fields

## Example

```go
package main

import (
  "log"
  "os"

  "github.com/miolini/cliconfig"
  "github.com/urfave/cli"
)

type Config struct {
  ListenAddr string `flag:"listen_addr_flag" env:"LISTEN_ADDR_ENV" default:"localhost:8080"`
}

func main() {
  config := Config{}
  app := cli.NewApp()
  app.Name = "example-app"
  app.Flags = cliconfig.Fill(&config, "EXAMPLE_APP_")
  app.Action = func(ctx *cli.Context) error {
    log.Printf("config: %#v", config)
    return nil
  }
  app.Run(os.Args)
}
```

Usage

```$ ./simple  --help
command-line-arguments
NAME:
   example-app - A new cli application

USAGE:
   simple [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --listen_addr_flag value  (default: "localhost:8080") [$EXAMPLE_APP_LISTEN_ADDR_ENV]
   --help, -h                show help
   --version, -v             print the version
