package main

import (
	"anketovac/poll"
	_ "anketovac/templates"

	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/joho/godotenv"
)

func Pool(dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbpool", dbpool)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var verify string
	err = dbpool.QueryRow(context.Background(), "select 'verify'").Scan(&verify)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()
	r.Use(Pool(dbpool))

	r.Static("/assets", "./assets")

	r.GET("/", poll.Home)
	r.GET("/:id", poll.Show)
	g_poll := r.Group("/poll")
	{
		g_poll.POST("/create", poll.Create)
		g_poll.POST("/delete", poll.Delete)
		g_poll.POST("/edit", poll.Edit)
	}
	r.Run("0.0.0.0:4000")
}
