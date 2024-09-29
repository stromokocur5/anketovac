package poll

import (
	"anketovac/templates"
	"log"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"context"
	"fmt"
	_ "log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/joho/godotenv"
)

type PType string

type Poll struct {
	id          string
	title       string
	description string
	ptype       PType
	created_at  time.Time
}

const (
	Multiple PType = "multiple"
	Ranking  PType = "ranking"
	Image    PType = "image"
)

type PollOption struct {
	poll_id string
	name    string
	votes   uint
}

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func All(c *gin.Context) {
	dbpool := c.MustGet("dbpool").(*pgxpool.Pool)

	query := "select * from polls_anketovac"
	rows, err := dbpool.Query(context.Background(), query)
	defer rows.Close()
	var polls []Poll
	_ = polls

	if err == pgx.ErrNoRows {
		x := templates.Layout("x", templates.Home("no poll"))
		render(c, 200, x)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var poll Poll
		err := rows.Scan(&poll.id, &poll.title, &poll.description, &poll.ptype, &poll.created_at)
		if err != nil {
			log.Println(err)
			return
		}
		polls = append(polls, poll)

	}
	x := templates.Layout("x", templates.Home("cx"))

	render(c, 200, x)
}

func Show(c *gin.Context) {
	dbpool := c.MustGet("dbpool").(*pgxpool.Pool)
	id := c.Param("id")

	query := fmt.Sprintf("select * from polls_anketovac where id = '%s'", id)
	var poll string
	err := dbpool.QueryRow(context.Background(), query).Scan(&poll)
	if err == pgx.ErrNoRows {
		x := templates.Layout("x", templates.Home("no poll"))
		render(c, 200, x)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	x := templates.Layout("x", templates.Home(poll))

	render(c, 200, x)
}

func Create(c *gin.Context) {}

func Delete(c *gin.Context) {}

func Edit(c *gin.Context) {}
