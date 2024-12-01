package server

import (
	"lupa/calba/budden/internal/handler"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/ping", handler.Ping)

	r.Run()
}
