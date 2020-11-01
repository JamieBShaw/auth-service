package main

import (
	"context"
	"database/sql"
	"flag"
	rd "github.com/JamieBShaw/auth-service/repository/redis"
	"github.com/JamieBShaw/auth-service/repository/sqlite"
	service "github.com/JamieBShaw/auth-service/service"
	internalhttp "github.com/JamieBShaw/auth-service/transport/http"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	driver = "sqlite3"
	db = "auth"
)

var (
	log            = logrus.New()
	router         = mux.NewRouter()
	port           = os.Getenv("PORT")
	sqlFlag = flag.Bool("sql", false, "service will use Sqlite3 as repo instead as default Redis")
)

func main() {
	if *sqlFlag {
		db, err := sql.Open(driver, db)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		sqlRepo := sqlite.NewRepository(db)
		authService := service.NewAuthService(sqlRepo, log)
		_ = internalhttp.NewHttpServer(authService, router, log)

	} else {

		dsn := os.Getenv("REDIS_DSN")
		if len(dsn) == 0 {
			dsn = "localhost:6379"
		}
		client := redis.NewClient(&redis.Options{
			Addr: dsn, //redis port
			Password: "",
			DB: 0,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Fatal(err)
		}
		redisRepo := rd.NewRepo(client, log)

		authService := service.NewAuthService(redisRepo, log)
		handler := internalhttp.NewHttpServer(authService, router, log)

		srv := &http.Server{
			Addr:         "0.0.0.0:" + port,
			Handler:      handler,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		go func() {
			log.Infof("Starting HTTP User Service running on port: %v", port)
			if err := srv.ListenAndServe(); err != nil {
				log.Fatal(err)
			}
		}()

		c := make(chan os.Signal, 1)
		// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
		// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
		signal.Notify(c, os.Interrupt)

		// Block until we receive our signal.
		<-c

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		srv.Shutdown(ctx)
		// Optionally, you could run srv.Shutdown in a goroutine and block on
		// <-ctx.Done() if your application should wait for other services
		// to finalize based on context cancellation.
		log.Println("shutting down")
		os.Exit(0)

	}




}
