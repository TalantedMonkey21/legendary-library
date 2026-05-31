package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TalantedMonkey21/GoLectures/internal/config"
	"github.com/TalantedMonkey21/GoLectures/internal/repository"
	httptransport "github.com/TalantedMonkey21/GoLectures/internal/transport/http"
	"github.com/TalantedMonkey21/GoLectures/internal/usecase"
	"gorm.io/gorm"
)


func doRetryConnection(cfg *config.Config) (*gorm.DB, error) {
	maxAttempts := 10
	for i := range(maxAttempts) {
		db, err := config.ConnectDB(cfg.DB)
		if err == nil {
			return db, nil
		}
		log.Printf("Error connecting to database: %v, retrying in 3 seconds...\n", err)
		if i < maxAttempts {
			time.Sleep(3 * time.Second)
		}
	}
	return nil, errors.New("failed to connect to database after 10 attempts")
}

func main() {
	// получаем конфиги
	cfg := config.Load()

	// TODO:
	// может случится такое, что сервер поднимается раньше через БД
	// и сейчас подключение происходит 1 раз
	// нужно подумать над тем как сделать 10 попыток каждые 3 секунды

	// создаем соединение с бд
	db, err := doRetryConnection(cfg)
	if err != nil {
		log.Println("can not connect to database:", err)
		os.Exit(1)
	}

	// автомиграция
	if err := config.AutoMigrate(db); err != nil {
		log.Println("can not migrate to database:", err)
		os.Exit(1)
	}

	// слой хранилища
	noteRepo := repository.NewNoteRepo(db)

	// слой бизнес логики
	noteService := usecase.New(noteRepo)

	// слой хэндлеров
	handler := httptransport.NewHandler(noteService)

	// роутер
	router := httptransport.NewRouter(handler)

	server := &http.Server{
		Addr: cfg.Port,
		Handler: router,
		ReadTimeout: 60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	go func() {
		log.Printf("HTTP started on port: %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server error: %v", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<- stop
	log.Println("Signal recived")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Forced shutdown: %v", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}
