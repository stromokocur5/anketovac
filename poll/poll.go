package poll

import (
	"github.com/a-h/templ"
	_ "github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.String(200, c.Param("poll"))
}

func Create(c *gin.Context) {}

func Delete(c *gin.Context) {}

func Edit(c *gin.Context) {}
