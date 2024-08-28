package handler

import (
	"ToDo-List/docs"
	"ToDo-List/pkg/models"
	"ToDo-List/pkg/service"
	"ToDo-List/pkg/storage"
	"errors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

const BasePath = "/"

// @BasePath /
func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = BasePath
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/tasks", h.CreateTask)
	r.GET("/tasks", h.GetAllTasks)
	r.GET("/tasks/:id", h.GetTaskByID)
	r.PUT("/tasks/:id", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)

	return r
}

// CreateTask godoc
// @Summary Создание новой задачи
// @Schemes
// @Description Создание новой задачи
// @Accept json
// @Produce json
// @Param data body models.CreateTaskReq true "Входные параметры"
// @Success 201 {object} models.Task
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /tasks [post]
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

// GetAllTasks godoc
// @Summary Получение списка всех задач
// @Schemes
// @Description Получение списка всех задач
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 500 {object} Error
// @Router /tasks [get]
func (h *Handler) GetAllTasks(c *gin.Context) {
	tasks, err := h.services.IToDoService.GetAllTasks()
	if err != nil {
		log.Println("Handler. GetAllTasks:", err)
		c.JSON(http.StatusInternalServerError, Error{Message: IternalServerErrorTitle})
		return
	}

	if len(tasks) == 0 {
		tasks = make([]models.Task, 0)
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID godoc
// @Summary Получение задачи по ID
// @Schemes
// @Description Получение задачи по ID
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Task
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /tasks/{id} [get]
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

// UpdateTask godoc
// @Summary Обновление задачи
// @Schemes
// @Description Обновление задачи
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param data body models.CreateTaskReq true "Входные параметры"
// @Success 200 {object} models.Task
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /tasks/{id} [put]
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

// DeleteTask godoc
// @Summary Удаление задачи
// @Schemes
// @Description Удаление задачи
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 204
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /tasks/{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: BadRequestTitle})
		return
	}

	err = h.services.IToDoService.DeleteTask(id)
	if err != nil {
		log.Println("Handler. DeleteTask:", err)
		switch {
		case errors.Is(err, storage.TaskNotFound):
			c.JSON(http.StatusNotFound, Error{Message: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{Message: IternalServerErrorTitle})
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)
}
