package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/okoge-kaz/golang-todo-application/server/controllers"
	"github.com/okoge-kaz/golang-todo-application/server/helpers"
)

// Init initializes the router
func Init() *gin.Engine {
	// initialize router
	router := gin.Default()

	// initialize session
	helpers.SetSession(router)

	// routing
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// task
	router.GET("/task", controllers.ShowTasks)
	task := router.Group("/task")
	// taskGroup.Use(service.LoginCheck) // login check
	{
		// create
		task.POST("/new", controllers.CreateTask)

		// read
		task.GET("/:id", controllers.ShowTask)

		// update
		taskUpdate := task.Group("/:id/")
		// taskUpdate.Use(service.TaskOwnerCheck) // task owner check ユーザーとの紐付け
		{
			// edit
			taskUpdate.POST("edit", controllers.UpdateTask)

			// delete
			taskUpdate.POST("delete", controllers.DeleteTask)
		}
	}

	// user
	// new registration
	router.POST("/user/new", controllers.CreateUser)

	user := router.Group("/user")
	// user.Use(service.LoginCheck) // login check
	{
		// read
		user.GET("/info", controllers.NotImplemented)

		// change password
		user.POST("/change_password", controllers.NotImplemented)

		// change user name
		user.POST("/change_name", controllers.NotImplemented)

		// delete
		user.POST("/delete", controllers.NotImplemented)
	}

	// category
	router.GET("/category", controllers.NotImplemented)
	category := router.Group("/category")
	{
		// create
		category.POST("/new", controllers.NotImplemented)
		// delete
		category.POST("/:id/delete", controllers.NotImplemented)
	}

	// login & logout
	router.POST("/login", controllers.NotImplemented)
	router.POST("/logout", controllers.NotImplemented)

	return router
}
