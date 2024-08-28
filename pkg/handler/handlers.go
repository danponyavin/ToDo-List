package handler

import (
	"ToDo-List/pkg/models"
	"ToDo-List/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	BadRequestTitle         = "Неправильный формат данных."
	IternalServerErrorTitle = "Проблема на сервере."
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

type Error struct {
	Message string `json:"message"`
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", h.CreateTask)
	r.GET("/tasks", h.GetAllTasks)
	r.GET("/tasks/{id}", h.GetTaskByID)
	r.PUT("/tasks/{id}", h.UpdateTask)
	r.DELETE("/tasks/{id}", h.DeleteTask)

	return r
}

func (h *Handler) CreateTask(c *gin.Context) {
	var req models.CreateTaskReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: BadRequestTitle})
		return
	}

	task, err := h.services.IToDoService.CreateTask(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: IternalServerErrorTitle})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *Handler) GetAllTasks(c *gin.Context) {}

func (h *Handler) GetTaskByID(c *gin.Context) {}

func (h *Handler) UpdateTask(c *gin.Context) {}

func (h *Handler) DeleteTask(c *gin.Context) {}
