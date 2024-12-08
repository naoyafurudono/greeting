package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/naoyafurudono/clio-go"
	"github.com/naoyafurudono/proto-cli/gen/greet/v1/greetv1clio"
	"github.com/naoyafurudono/proto-cli/service"
)

// The entry point (what you implement)
func main() {
	greetCmd := greetv1clio.NewGreetServiceCommand(context.Background(), &service.GreetServer{}, os.Stdout)
	if err := greetCmd.Execute(); err != nil {
		if errors.Is(err, clio.CLIFailed) {
			panic(err)
		} else if errors.Is(err, clio.RPCFailed) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			panic("never")
		}
	}
}
