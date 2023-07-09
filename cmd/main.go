package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	manage "soat1-challenge1/internal"
	"soat1-challenge1/internal/core/usecases"
	repository "soat1-challenge1/internal/repositories"

	"github.com/gin-gonic/gin"
)

var httpPort = fmt.Sprintf(":%s", os.Getenv("API_PORT"))

const (
	shutdownTimeout = 5 * time.Second
	pathPrefix      = "/api/v1"
)

func main() {
	fmt.Println("port", httpPort)
	repository := repository.NewRepository()

	// create manager
	m := manage.New(&manage.UseCases{
		Order:           usecases.NewOrderUseCase(repository.Order),
		Product:         usecases.NewProductUseCase(repository.Product),
		ProductCategory: usecases.NewProductCategoryUseCase(repository.ProductCategory),
		Customer:        usecases.NewCustomerUseCase(repository.Customer),
	})

	// create engine
	engine := gin.Default()

	// base route
	v1Routes := engine.Group(pathPrefix)

	// routes
	m.RegisterRoutes(v1Routes)

	// Inicia server
	engine.Run(httpPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: engine,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
	}
}
