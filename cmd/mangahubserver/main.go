package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func run(ctx context.Context, name string, args ...string) {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Starting MangaHub services...")

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go run(ctx, "go", "run", "../api_server/main.go")
	go run(ctx, "go", "run", "../udp_server/main.go")
	go run(ctx, "go", "run", "../grpc_server/main.go")
	go run(ctx, "go", "run", "../ws_server/main.go")

	<-ctx.Done()

	log.Println("Shutting down all services...")
}
