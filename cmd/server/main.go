package main

import (
	"context"
	"fmt"
	"os"

	"github.com/naoyafurudono/proto-cli/my_gen"
	"github.com/naoyafurudono/proto-cli/service"
)

// The entry point (what you implement)
func main() {
	var greetCmd = my_gen.NewGreetCommand(context.Background(), &service.GreetServer{})
	if err := greetCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
