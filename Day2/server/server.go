package server

import (
	"SofwareGoDay2/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func NewServer() (*Server) {
	r := gin.Default()
	routes.ApplyRoutes(r)
	return &Server{r}
}