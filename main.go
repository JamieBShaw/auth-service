package main

import (
	"context"
	"flag"
	"github.com/JamieBShaw/auth-service/protob"
	rd "github.com/JamieBShaw/auth-service/repository/redis"
	"github.com/JamieBShaw/auth-service/service"
	internalGrpc "github.com/JamieBShaw/auth-service/transport/grpc"
	internalhttp "github.com/JamieBShaw/auth-service/transport/http"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	googlegrpc "google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	log    = logrus.New()
	router = mux.NewRouter()
	port   = os.Getenv("PORT")
	grpc   = flag.Bool("grpc", true, "service will use grpc (http2) as the transport layer")
	//sqlFlag = flag.Bool("sql", false, "service will use Sqlite3 as repo instead as default Redis")
)

func main() {

	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "0.0.0.0:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     dsn, //redis port
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	redisRepo := rd.NewRepo(client, log)
	authService := service.NewAuthService(redisRepo, log)

	if *grpc {
		log.Infof("Starting GRPC Auth Service running on port: %v", port)

		lis, err := net.Listen("tcp", "0.0.0.0:"+port)
		if err != nil {
			log.Fatal("Failed to listen", err)
		}

		s := googlegrpc.NewServer()
		srv := internalGrpc.NewGrpcServer(authService)
		protob.RegisterAuthServiceServer(s, srv)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	} else {
		handler := internalhttp.NewHttpServer(authService, router, log)

		srv := &http.Server{
			Addr:         "0.0.0.0:" + port,
			Handler:      handler,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		go func() {
			log.Infof("Starting HTTP Auth Service running on port: %v", port)
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
