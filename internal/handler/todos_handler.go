package handler

import (
	"lupa/calba/budden/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Repo *repository.TodoRepository
}

func (h *TodoHandler) ListTodos(c *gin.Context) {
	todos, err := h.Repo.ListTodos(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list todos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var newTodo repository.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.CreateTodo(c.Request.Context(), newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"todo": newTodo})
}
