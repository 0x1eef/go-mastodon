package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestCmdReblog(t *testing.T) {
	ok := false
	f := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v1/statuses/123/reblog":
			fmt.Fprintln(w, `{}`)
			ok = true
			return
		}
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	testWithServer(
		f, func(app *cli.App) {
			app.Run([]string{"mstdn", "reblog", "122"})
		},
	)
	if ok {
		t.Fatal("expected 'not found' response")
	}

	ok = false
	testWithServer(
		f, func(app *cli.App) {
			app.Run([]string{"mstdn", "reblog", "123"})
		},
	)
	if !ok {
		t.Fatal("expected an 'OK' response")
	}
}
