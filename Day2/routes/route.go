package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func world(c *gin.Context){
	message := os.Getenv("HELLO_MESSAGE")
	if message == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "%s", message)
}

func health(c *gin.Context) {
	c.Status(http.StatusOK)
}

func query(c *gin.Context){
	Message := c.Query("message")
	if Message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "Query is %s", Message)
}

func param(c *gin.Context){
	Message := c.Param("message")
	if Message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "Param is :'%s'", Message)
}

func body(c *gin.Context){
	Message := c.PostForm("message")
	if Message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "Body value is '%s'", Message)
}

func header(c *gin.Context){
	Message := c.GetHeader("message")
	if Message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "Header value:'%s'", Message)
}

func cookie(c *gin.Context){
	cookie, err := c.Cookie("message")

	if err != nil || cookie == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Printf("Cookie value: %s \n", cookie)
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/health", health)
	r.GET("/repeat-my-query", query)
	r.GET("/repeat-my-param/:message", param)
	r.POST("/repeat-my-body", body)
	r.GET("/repeat-my-header", header)
	r.GET("/repeat-my-cookie", cookie)
}
