package poll

import (
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.String(200, c.Param("posll"))
}

func Create(c *gin.Context) {}

func Delete(c *gin.Context) {}

func Edit(c *gin.Context) {}
