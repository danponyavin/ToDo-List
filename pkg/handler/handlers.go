package handler

import (
	"ToDo-List/pkg/models"
	"ToDo-List/pkg/service"
	"ToDo-List/pkg/storage"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	r.GET("/tasks/:id", h.GetTaskByID)
	r.PUT("/tasks/:id", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)

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
		log.Println("Handler. CreateTask:", err)
		c.JSON(http.StatusInternalServerError, Error{Message: IternalServerErrorTitle})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *Handler) GetAllTasks(c *gin.Context) {
	tasks, err := h.services.IToDoService.GetAllTasks()
	if err != nil {
		log.Println("Handler. GetAllTasks:", err)
		c.JSON(http.StatusInternalServerError, Error{Message: IternalServerErrorTitle})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTaskByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: BadRequestTitle})
		return
	}

	task, err := h.services.IToDoService.GetTask(id)
	if err != nil {
		log.Println("Handler. GetTaskByID:", err)
		switch {
		case errors.Is(err, storage.TaskNotFound):
			c.JSON(http.StatusNotFound, Error{Message: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{Message: IternalServerErrorTitle})
			return
		}
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: BadRequestTitle})
		return
	}

	var req models.CreateTaskReq
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: BadRequestTitle})
		return
	}

	task, err := h.services.IToDoService.UpdateTask(id, req)
	if err != nil {
		log.Println("Handler. UpdateTask:", err)
		switch {
		case errors.Is(err, storage.TaskNotFound):
			c.JSON(http.StatusNotFound, Error{Message: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{Message: IternalServerErrorTitle})
			return
		}
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTask(c *gin.Context) {}
