package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/okoge-kaz/golang-todo-application/server/controllers"
	"github.com/okoge-kaz/golang-todo-application/server/controllers/auth"
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
	task.Use(auth.LoginCheck)
	{
		// create
		task.POST("/new", controllers.CreateTask)

		// read
		task.GET("/:id", controllers.ShowTask)

		// update
		taskUpdate := task.Group("/:id/")
		taskUpdate.Use(auth.TaskOwnerCheck) // task owner check
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
	user.Use(auth.LoginCheck)
	{
		// read (user information)
		user.GET("/info", controllers.ShowUser)

		// change password
		user.POST("/change_password", controllers.ChangeUserPassword)

		// change user name
		user.POST("/change_name", controllers.ChangeUserName)

		// delete
		user.POST("/delete", controllers.DeleteUser)
	}

	// category
	router.GET("/category", controllers.NotImplemented)
	category := router.Group("/category")
	category.Use(auth.LoginCheck)
	{
		// create
		category.POST("/new", controllers.NotImplemented)
		// delete
		category.POST("/:id/delete", controllers.NotImplemented)
	}

	// login & logout
	router.POST("/login", controllers.NotImplemented)
	router.POST("/logout", controllers.NotImplemented)

	// login check
	router.GET("/login_check", controllers.NotImplemented)

	return router
}
