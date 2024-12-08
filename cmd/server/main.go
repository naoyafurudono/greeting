package main

import (
	"context"
	"fmt"
	"os"

	"github.com/naoyafurudono/proto-cli/gen/greet/v1/greetv1clio"
	"github.com/naoyafurudono/proto-cli/service"
)

// The entry point (what you implement)
func main() {
	greetCmd := greetv1clio.NewGreetServiceCommand(context.Background() , &service.GreetServer{})
	if err := greetCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
