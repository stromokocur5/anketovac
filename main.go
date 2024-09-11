package main

import (
	"anketovac/poll"
	_ "anketovac/templates"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:poll", poll.Show)
	g_poll := r.Group("/poll")
	{
		g_poll.GET("/create", poll.Create)
		g_poll.GET("/delete", poll.Delete)
		g_poll.GET("/edit", poll.Edit)
	}
	r.Run("0.0.0.0:4000")
}
