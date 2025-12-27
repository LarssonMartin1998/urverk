package main

import "urverk-backend/internal/httpserver"

func main() {
	httpserver.ListenAndServeAndAwaitGracefulShutdown(httpserver.CreateRoutes())
}
