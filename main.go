package main

import (
	"fmt"

	"github.com/alexander-littleton/go-htmx-project/pkg/db"
	"github.com/alexander-littleton/go-htmx-project/pkg/webserver"
)

func main() {

	conn, err := db.ConnectToDb()
	if err != nil {
		fmt.Printf("db connection failed: %s", err.Error())
		return
	}
	webserver.Init(conn)
}
