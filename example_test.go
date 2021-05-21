package graceful_test

import (
	"fmt"
	"net/http"

	"github.com/hyqe/graceful"
)

func Example() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})

	server := graceful.NewServer(
		graceful.WithHandler(handler),
		graceful.WithPort(8080),
	)

	graceful.Run(server)
}
