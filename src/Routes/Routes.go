
package Routes

import (
	"Controllers"
	"github.com/gin-gonic/gin"
)

/*
func LoggingMiddleware(c *gin.Context) {

	reqElements := make(map[string]interface{})
	reqElements["method"] = c.Request.Method
	println("Kostas: " + c.Request.Method)
	c.Next()
}*/

func SetupRouter() *gin.Engine {

	r := gin.Default()
    
	r.LoadHTMLGlob("templates/**/*")
    r.Static("/public","./public")

	v1 := r.Group("/tasks")
	{
		v1.GET("list", Controllers.GetList)
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}
	return r
}