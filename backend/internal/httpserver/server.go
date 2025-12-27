// Package httpserver
package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func ListenAndServeAndAwaitGracefulShutdown(handler http.Handler) error {
	server := createServerFromConfig(handler)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return onServerShuttingDown(server)
}

func onServerShuttingDown(server *http.Server) error {
	log.Println("🛑 Shutting down server...")
	// ctx, cancel := context.WithTimeout(context.Background(), max(cfg.Server.ReadTimeout, cfg.Server.WriteTimeout))
	ctx, cancel := context.WithTimeout(context.Background(), max(1, 5)) // Temp values until we have a config setup
	defer cancel()

	// TODO: Invoke shutdown server callbackd
	// perfect for DB to listen to and shutdown from etc.
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("❌ Server forced to shutdown: %v", err)
		return err
	}

	log.Println("✅ Server gracefully stopped")
	return nil
}

func createServerFromConfig(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":3000",
		Handler: handler,
		// TODO: Fix config
		// ReadTimeout:    cfg.Server.ReadTimeout,
		// WriteTimeout:   cfg.Server.WriteTimeout,
		// IdleTimeout:    cfg.Server.IdleTimeout,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
}
