package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	ctx context.Context
}


func (server *Server) NewServer(ctx context.Context) {
	server.ctx = ctx
	router := gin.Default()
	
	router.POST("/create", server.create)
	router.GET("/get/:id", server.get)
	router.UPDATE("/update/:id", server.update)
	router.DELETE("/delete/:id", server.delete)

	server.router = router
	return server
}

func (server *Server) StartServer(port string) error {
	return server.router.Run(port)
}