package Controllers

import (
	"net/http"
	"Models"
	"github.com/gin-gonic/gin"
)


func GetList(c *gin.Context) {
	var todo []Models.Todo

	err := Models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		
	    c.JSON(http.StatusOK, todo)
	
	//	c.HTML(http.StatusOK,"items.html",gin.H{
	//		"todo":todo,
	//		"tasks":"Λίστα Καθηκόντων",
	//		})
	}
}


func GetTodos(c *gin.Context) {
    
    	c.HTML(http.StatusOK,"items.html",gin.H{
			"tasks":"Λίστα Καθηκόντων",
			})
}


func CreateATodo(c *gin.Context) {
	var todo Models.Todo
	c.BindJSON(&todo)
	err := Models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodo(c *gin.Context) {
 
    id := c.Params.ByName("id")
    var todo Models.Todo
	err := Models.GetATodo(&todo, id)
	
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = Models.UpdateATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Models.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
