package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qytela/example-project-layout/internal/app/http"
	"github.com/qytela/example-project-layout/internal/app/providers"
	"github.com/qytela/example-project-layout/internal/pkg/exception"
	"github.com/qytela/example-project-layout/internal/pkg/logger"
	"github.com/qytela/example-project-layout/internal/pkg/validation"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.MakeLogEntry(nil).Panic("failed to load .env: ", err)
	}

	// Call the providers
	providers.ProvideDB()
	providers.ProvideSupabase()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.HTTPErrorHandler = exception.HTTPErrorHandler
	e.Validator = validation.NewValidation()

	handler := http.NewHandler()
	http.NewServer(e, handler)

	address := "127.0.0.1:8080"
	logger.MakeLogEntry(nil).Info("server started on address: ", address)

	if err := e.Start(address); err != nil {
		logger.MakeLogEntry(nil).Panic(err)
	}
}
