package main

import (
	"log"
	"net/http"
	"os"
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
	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	fmt.Printf("Error opening DB: %v", err)
	// }

	// redi := redis.NewClient(&redis.Options{
	// 	Addr: redisAddr,
	// })
	//sessionStore := sessions.NewRedisStore(redi, time.Hour)
	//sql := &users.MySQLStore{Database: db}

	mux := http.NewServeMux()

	log.Printf("Server is listening at %s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
}