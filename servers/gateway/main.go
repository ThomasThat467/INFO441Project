package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ThomasThat467/INFO441Project/tree/main/servers/handlers"
	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/plants"
	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/schedules"
	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/users"
	"github.com/ThomasThat467/INFO441Project/tree/main/servers/sessions"
	"github.com/go-redis/redis"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")

	sessionKey := os.Getenv("SESSIONKEY")
	if len(sessionKey) == 0 {
		log.Fatalf("SESSIONKEY is not set")
	}

	redisAddr := os.Getenv("REDISADDR")
	if len(redisAddr) == 0 {
		log.Fatalf("REDDISADDR is not set")
	}
	dsn := os.Getenv("DSN")
	if len(dsn) == 0 {
		log.Fatalf("DSN is not set")
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error opening DB: %v", err)
	}

	redi := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	sessionStore := sessions.NewRedisStore(redi, time.Hour)
	userStore := &users.MySQLStore{Database: db}
	plantStore := &plants.MySQLStore{Database: db}
	scheduleStore := &schedules.MySQLStore{Database: db}
	ctx := handlers.NewHandlerContext(sessionKey, sessionStore, userStore, plantStore, scheduleStore)

	mux := http.NewServeMux()
	corsmux := handlers.NewCorsHandler(mux)

	mux.HandleFunc("/v1/users", ctx.UsersHandler)
	mux.HandleFunc("/v1/users/", ctx.SpecificUserHandler)
	mux.HandleFunc("/v1/sessions", ctx.SessionsHandler)
	mux.HandleFunc("/v1/sessions/", ctx.SpecificSessionHandler)
	mux.HandleFunc("/v1/plant", ctx.PlantHandler)
	mux.HandleFunc("/v1/schedule", ctx.ScheduleHandler)

	log.Printf("Server is listening at %s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, corsmux))
}
