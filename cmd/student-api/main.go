package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZeeshanSaleem-official/student-api/internal/config"
	"github.com/ZeeshanSaleem-official/student-api/internal/http/handlers/student"
	"github.com/ZeeshanSaleem-official/student-api/internal/storage/postgresql"
)

func main() {
	// config file
	cfg := config.MustLoad()
	// database setup

	storage, err := postgresql.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Database is initialized", slog.String("env", cfg.Env))

	// router setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	router.HandleFunc("PUT /api/students/{id}", student.Update(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.Delete(storage))

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started ", slog.String("at address: ", cfg.Addr))

	// Gracefully Server Shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()
	<-done

	slog.Info("Shutting Down the server!!")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to Shutdown the server:", slog.String("Error: ", err.Error()))
	}
	slog.Info("Server Shutdown successfully!!")

}
