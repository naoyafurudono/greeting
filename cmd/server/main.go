package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/naoyafurudono/clio-go"
	"github.com/naoyafurudono/proto-cli/gen/greet/v1/greetv1clio"
	"github.com/naoyafurudono/proto-cli/gen/greet/v1/greetv1connect"
	"github.com/naoyafurudono/proto-cli/service"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	if err := root.Execute(); err != nil {
		if errors.Is(err, clio.RPCFailed) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else if errors.Is(err, clio.CLIFailed) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		} else {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(3)
		}
	}
}

var root = cobra.Command{
	Use:   "greeting",
	Short: "greeting",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	// 普通のconnect serverと同居させてみる
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "connect server",
		RunE: func(cmd *cobra.Command, args []string) error {
			greeter := &service.GreetServer{}
			mux := http.NewServeMux()
			path, handler := greetv1connect.NewGreetServiceHandler(greeter)
			mux.Handle(path, handler)
			return http.ListenAndServe(
				"localhost:8080",
				// Use h2c so we can serve HTTP/2 without TLS.
				h2c.NewHandler(mux, &http2.Server{}),
			)
		},
	}
	greetCmd := greetv1clio.NewGreetServiceCommand(context.Background(), &service.GreetServer{}, os.Stdout)
	root.AddCommand(serveCmd, greetCmd)
}
