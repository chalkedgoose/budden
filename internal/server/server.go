package server

import (
	"fmt"
	"lupa/calba/budden/internal/config"
	"lupa/calba/budden/internal/handler"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.LoadConfig()
	r := gin.Default()

	r.GET("/ping", handler.Ping)

	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
