package handler

import (
	"ToDo-List/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/tasks", h.CreateTask)
	r.GET("/tasks", h.GetAllTasks)
	r.GET("/tasks/{id}", h.GetTaskByID)
	r.PUT("/tasks/{id}", h.UpdateTask)
	r.DELETE("/tasks/{id}", h.DeleteTask)

	return r
}

func (h *Handler) CreateTask(c *gin.Context) {}

func (h *Handler) GetAllTasks(c *gin.Context) {}

func (h *Handler) GetTaskByID(c *gin.Context) {}

func (h *Handler) UpdateTask(c *gin.Context) {}

func (h *Handler) DeleteTask(c *gin.Context) {}
