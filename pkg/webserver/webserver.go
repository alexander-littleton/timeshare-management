package webserver

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/alexander-littleton/go-htmx-project/pkg/domain/calendar"
	"github.com/alexander-littleton/go-htmx-project/pkg/domain/user"
	"github.com/jackc/pgx/v5"
)

func Init(dbConn *pgx.Conn) {
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(Home()))

	InitCalendarRoutes(mux)
	InitUserRoutes(mux, dbConn)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", mux)
}

func InitCalendarRoutes(mux *http.ServeMux) {
	calendarSvc := calendar.NewService()
	controller := calendar.NewController(calendarSvc)

	mux.HandleFunc("GET /calendar", controller.GetCalendar)
}

func InitUserRoutes(mux *http.ServeMux, dbConn *pgx.Conn) {
	repo := user.NewRepo(dbConn)
	userService := user.NewService(repo)
	controller := user.NewController(userService)

	mux.HandleFunc("/user", controller.CreateUser)
}
