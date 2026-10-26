package app

import (
	"context"
	"log"
	"os/signal"
	"portalnews/config"
	"portalnews/lib/auth"
	"portalnews/lib/middleware"
	"portalnews/lib/pagination"
	"syscall"
	"time"

	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
)

func RunServer() {
	cfg := config.NewConfig()

	_, err := cfg.ConnectPostgres()
	if err != nil {
		log.Fatalf("error to connect database %v", err)
		return
	}

	// cloudflare
	cdfR2 := cfg.LoadAwsConfig()
	_ = s3.NewFromConfig(cdfR2)

	// jwt
	_ = auth.NewJwt(cfg)

	// middleware
	_ = middleware.NewMiddleware(cfg)

	_ = pagination.NewPagination()

	app := fiber.New()       // instance fiber
	app.Use(cors.New())      // middleware cors
	app.Use(recoverer.New()) // middleware handle panic()
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${status} - ${method} ${path}\n",
	})) // mencatat informasi request yang masuk ex. [2026-10-26 09:30:00] 127.0.0.1:54321 200 - GET /api...

	_ = app.Group("/api") // mengelompokkan endpoint dengan prefix tertentu

	// go routine
	go func() {
		if cfg.App.AppPort == "" {
			cfg.App.AppPort = os.Getenv("APP_PORT")
		}
		err := app.Listen(":" + cfg.App.AppPort)
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Menangani signal untuk shutdown
	quit := make(chan os.Signal, 1)      // membuat channel
	signal.Notify(quit, os.Interrupt)    // mendaftarkan sinyal user menekan Ctrl + C
	signal.Notify(quit, syscall.SIGTERM) // meminta proses berhenti secara normal
	<-quit

	log.Println("server shutdown in 5 second")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // memberikan batas waktu maksimal untuk proses shutdown

	defer cancel() // melepaskan resource yang berkaitan dengan context

	app.ShutdownWithContext(ctx) // Fiber menghentikan server dengan context yang sudah dibuat
}
