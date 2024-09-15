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
	// defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'verify'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()
	r.Use(Pool(dbpool))

	r.GET("/:poll", poll.Show)
	g_poll := r.Group("/poll")
	{
		g_poll.GET("/create", poll.Create)
		g_poll.GET("/delete", poll.Delete)
		g_poll.GET("/edit", poll.Edit)
	}
	r.Run("0.0.0.0:4000")
}
