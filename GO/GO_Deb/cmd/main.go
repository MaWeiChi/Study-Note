package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const APIPath = "/api"

var VERSION string

type Event interface {
	Init() error
	PostEvents(*gin.Context)
	GetEvents(*gin.Context)
}

type event struct {
	Origin     string `json:"origin"`
	Severity   string `json:"severity"`
	Category   string `json:"category"`
	Event      string `json:"event"`
	Message    string `json:"message"`
	CreatedAt  string `json:"created"`
	User       string `json:"user"`
	UserOrigin string `json:"userOrigin"`
}

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	router := gin.Default()
	router.Use()
	srv := &http.Server{
		Addr:    ":59000",
		Handler: router,
	}

	router.GET(APIPath+"/demo", Get)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	waitTimer, terminal := context.WithTimeout(context.Background(), 10*time.Second)
	defer terminal()
	if err := srv.Shutdown(waitTimer); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "demo"})
}
