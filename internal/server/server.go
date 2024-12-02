package server

import (
	"fmt"
	"lupa/calba/budden/internal/config"
	"lupa/calba/budden/internal/handler"
	"lupa/calba/budden/internal/repository"

	"github.com/gin-gonic/gin"
)

func Run() {
	todoRepo, err := repository.NewTodoRepository()
	if err != nil {
		return
	}
	defer todoRepo.Close()

	cfg := config.LoadConfig()
	r := gin.Default()

	todoHandler := &handler.TodoHandler{Repo: todoRepo}

	r.GET("/ping", handler.Ping)
	r.GET("/todos", todoHandler.ListTodos)
	r.POST("/todos", todoHandler.CreateTodo)

	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
