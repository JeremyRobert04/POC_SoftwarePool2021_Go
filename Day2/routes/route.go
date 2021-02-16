package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func world(c *gin.Context){
	c.JSON(http.StatusOK, "world")
}

func query(c *gin.Context){
	Message := c.Query("message")
	if Message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.String(http.StatusOK, "Query is %s", Message)
}

func param(c *gin.Context){
	Message := c.Param("message")
	if Message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.String(http.StatusOK, "Param is :'%s'", Message)
}

func body(c *gin.Context){
	Message := c.PostForm("message")
	if Message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.String(http.StatusOK, "Body value is '%s'", Message)
}

func header(c *gin.Context){
	Message := c.GetHeader("message")
	if Message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.String(http.StatusOK, "Header value:'%s'", Message)
}

func cookie(c *gin.Context){
	cookie, err := c.Cookie("message")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("message", "test", 3600, "/", "localhost", false, true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/repeat-my-query", query)
	r.GET("/repeat-my-param/:message", param)
	r.POST("/repeat-my-body", body)
	r.GET("/repeat-my-header", header)
	r.GET("/repeat-my-cookie", cookie)
}
