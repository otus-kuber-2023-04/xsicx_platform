package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {

	apiServer := fiber.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	serverShutdown := make(chan struct{})

	go func() {
		<-c

		log.Print("Gracefully shutting down...")

		shutdownErr := apiServer.Shutdown()

		must(shutdownErr)

		log.Print("API server stopped")

		serverShutdown <- struct{}{}
	}()

	apiServer.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})
	apiServer.Static("/", "./app")

	must(apiServer.Listen(":8000"))

	<-serverShutdown
}

func must(err error) {
	if err == nil {
		return
	}

	log.Panic(err)
}
