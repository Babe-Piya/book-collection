package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github/Babe-piya/book-collection/appconfig"
	"github/Babe-piya/book-collection/database"

	"github.com/gin-gonic/gin"
)

func Start(config *appconfig.AppConfig) (*http.Server, *sql.DB) {
	gormDB, db, err := database.NewConnection(config.Database)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	Routes(router, gormDB, config)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.ServerPort),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
			log.Fatal("shutting down the server")
		}
	}()

	return srv, db
}

func Shutdown(serv *http.Server, db *sql.DB) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := serv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	if err := db.Close(); err != nil {
		log.Fatal("Failed to close DB:", err)
	}

	slog.Info("Shutdown complete")
}
