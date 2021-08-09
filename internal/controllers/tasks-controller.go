package controllers

import (
	"ai4am.com/go-restapi/internal/repository"
	"ai4am.com/go-restapi/models"
	"ai4am.com/go-restapi/pkg/helpers"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetTaskById godoc
// @Tags Tasks
// @Summary Retrieves task based on given ID
// @Description get Task by ID
// @Produce json
// @Param id path integer true "Task ID"
// @Success 200 {object} models.Task
// @Router /api/tasks/{id} [get]
// @Security Authorization Token
func GetTaskById(c *gin.Context) {
	s := repository.GetTaskRepository()
	id := c.Param("id")
	if task, err := s.GetTaskByID(id); err != nil {
		helpers.NewError(c, http.StatusNotFound, errors.New("task not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

// GetTasks godoc
// @Tags Tasks
// @Summary Retrieves tasks based on query
// @Description Get Tasks
// @Produce json
// @Param name query string false "Name"
// @Param text query string false "Text"
// @Success 200 {array} []models.Task
// @Router /api/tasks [get]
func GetTasks(c *gin.Context) {
	s := repository.GetTaskRepository()
	var q models.Task
	_ = c.Bind(&q)
	if tasks, err := s.QueryTasks(&q); err != nil {
		c.JSON(http.StatusOK, []models.Task{})
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

// GetTasks godoc
// @Tags Tasks
// @Summary Creates a new Task
// @Description Add Task
// @Accept  json
// @Produce json
// @Param user body models.Task true "Add task"
// @Success 201 {object} models.Task
// @Router /api/tasks [post]
// @Security Authorization Token
func CreateTask(c *gin.Context) {
	s := repository.GetTaskRepository()
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	fmt.Println(taskInput)
	if err := s.AddTask(&taskInput); err != nil {
		helpers.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, taskInput)
	}
}


// GetTasks godoc
// @Tags Tasks
// @Summary Updates an Task by ID
// @Description Update Task
// @Accept  json
// @Produce json
// @Param id path integer true "User ID"
// @Param user body models.Task true "Add task"
// @Success 200 {object} models.Task
// @Router /api/tasks/{id} [put]
// @Security Authorization Token
func UpdateTask(c *gin.Context) {
	s := repository.GetTaskRepository()
	id := c.Params.ByName("id")
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	if _, err := s.GetTaskByID(id); err != nil {
		helpers.NewError(c, http.StatusNotFound, errors.New("task not found"))
		log.Println(err)
	} else {
		if err := s.UpdateTask(&taskInput); err != nil {
			helpers.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, taskInput)
		}
	}
}

// GetTaskById godoc
// @Tags Tasks
// @Summary Deletes task based on given ID
// @Description Delete task
// @Produce json
// @Param id path integer true "Task ID"
// @Success 204
// @Router /api/task/{id} [get]
// @Security Authorization Token
func DeleteTask(c *gin.Context) {
	s := repository.GetTaskRepository()
	id := c.Params.ByName("id")
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	if task, err := s.GetTaskByID(id); err != nil {
		helpers.NewError(c, http.StatusNotFound, errors.New("task not found"))
		log.Println(err)
	} else {
		if err := s.Delete(task); err != nil {
			helpers.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
