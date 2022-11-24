package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/okoge-kaz/golang-todo-application/server/controllers"
)

// Init initializes the router
func Init() *gin.Engine {
	// initialize router
	router := gin.Default()

	// initialize session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

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
			taskUpdate.GET("edit", controllers.NotImplemented)
			taskUpdate.POST("edit", controllers.NotImplemented)

			// delete
			taskUpdate.POST("delete", controllers.NotImplemented)
		}
	}

	// user
	// new registration
	router.GET("/user/new", controllers.NotImplemented)
	router.POST("/user/new", controllers.NotImplemented)

	user := router.Group("/user")
	// user.Use(service.LoginCheck) // login check
	{
		// read
		user.GET("/info", controllers.NotImplemented)

		// change password
		user.GET("/change_password", controllers.NotImplemented)
		user.POST("/change_password", controllers.NotImplemented)

		// change user name
		user.GET("/change_name", controllers.NotImplemented)
		user.POST("/change_name", controllers.NotImplemented)

		// delete
		user.POST("/delete", controllers.NotImplemented)
	}

	// login & logout
	router.GET("/login", controllers.NotImplemented)
	router.POST("/login", controllers.NotImplemented)
	router.POST("/logout", controllers.NotImplemented)

	return router
}
