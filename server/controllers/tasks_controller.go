package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/okoge-kaz/golang-todo-application/server/entities"
	"github.com/okoge-kaz/golang-todo-application/server/helpers"
	"github.com/okoge-kaz/golang-todo-application/server/models"
)

func ShowTask(ctx *gin.Context) {
	// get task
	taskID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid task id)")
		return
	}

	task, err := models.GetTask(taskID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// render
	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func ShowTasks(ctx *gin.Context) {
	// get user_id from session
	userID := helpers.GetUserID(ctx)

	// query parameter
	keywords := strings.Split(ctx.Query("keywords"), " ") // "keyword1 keyword2 keyword3" -> ["keyword1", "keyword2", "keyword3"]
	order := ctx.Query("order")                           // "asc" or "desc" (if empty, "asc" is set)
	status := strings.Split(ctx.Query("status"), ",")     // "todo,in-progress,done" -> ["todo", "in-progress", "done"]

	var tasks []entities.Task
	var err error

	// status = [] :=> status = ["todo", "in-progress", "done"]
	if len(status) == 0 {
		status = []string{"todo", "in-progress", "done"}
	}
	// order = "" :=> order = "asc"
	if order == "" {
		order = "asc"
	}

	switch {
	case len(status) == 3 && len(keywords) == 0:
		// get all tasks
		tasks, err = models.GetTasksByUserID(userID, order)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	default:
		// get tasks by status and keywords
		tasks, err = models.GetTasksByUserIDAndStatusAndKeyword(userID, status, keywords, order)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	// render
	ctx.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func CreateTask(ctx *gin.Context) {
	// get user_id from session
	userID := helpers.GetUserID(ctx)

	var user entities.User
	var err error

	// get user by users.id
	if user, err = models.GetUserByID(userID); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// bind request body
	var task entities.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid request body)")
		return
	}

	// create task record
	if err := models.CreateTask(&task); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// create ownership record
	if err := models.CreateOwnership(user, task); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// render
	ctx.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}

func UpdateTask(ctx *gin.Context) {
	// get task
	taskID, err := strconv.Atoi(ctx.Param("task_id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid task id)")
		return
	}

	task, err := models.GetTask(taskID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// bind request body
	var updateTask entities.Task
	if err := ctx.BindJSON(&updateTask); err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid request body)")
		return
	}

	// update task record
	if err := models.UpdateTask(task, updateTask); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// render
	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func DeleteTask(ctx *gin.Context) {
	// get task
	taskID, err := strconv.Atoi(ctx.Param("task_id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid task id)")
		return
	}

	task, err := models.GetTask(taskID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// delete task record
	if err := models.DeleteTask(int(task.ID)); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// render
	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}
