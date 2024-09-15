package poll

import (
	_ "github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"context"
	"fmt"
	_ "log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/joho/godotenv"
)

func Show(c *gin.Context) {
	dbpool := c.MustGet("dbpool").(*pgxpool.Pool)
	var greeting string
	err := dbpool.QueryRow(context.Background(), "select 'verify'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	c.String(200, greeting)
}

func Create(c *gin.Context) {}

func Delete(c *gin.Context) {}

func Edit(c *gin.Context) {}
