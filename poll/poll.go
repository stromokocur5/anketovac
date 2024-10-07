package poll

import (
	. "anketovac/models"
	"anketovac/templates"
	"log"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"context"
	_ "fmt"
	_ "log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/joho/godotenv"
)

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func Home(c *gin.Context) {
	x := templates.Layout("x", templates.Home())

	render(c, 200, x)
}
func CreateView(c *gin.Context) {
	x := templates.Layout("x", templates.Home())

	render(c, 200, x)
}
func Show(c *gin.Context) {
	dbpool := c.MustGet("dbpool").(*pgxpool.Pool)
	id := c.Param("id")

	query := `
	select * from polls_anketovac where id = $1;
	`
	poll := Poll{}
	err := dbpool.QueryRow(context.Background(), query, id).Scan(&poll.Id, &poll.Title, &poll.Description, &poll.Ptype, &poll.Created_at)
	if err == pgx.ErrNoRows {
		x := templates.Layout("x", templates.PollView(Poll{}, []PollOption{}))
		render(c, 200, x)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}

	query = `
	select option_order,name,votes from poll_options where id = $1;
	`
	poll_options := []PollOption{}
	rows, err := dbpool.Query(context.Background(), query, id)

	defer rows.Close()

	for rows.Next() {
		poll_option := PollOption{}
		err = rows.Scan(&poll_option.OptionOrder, &poll_option.Name, &poll_option.Votes)
		if err != nil {
			log.Println(err)
			return
		}
		poll_options = append(poll_options, poll_option)
	}

	_ = poll_options
	x := templates.Layout("x", templates.PollView(poll, poll_options))

	render(c, 200, x)
}

func Create(c *gin.Context) {
	dbpool := c.MustGet("dbpool").(*pgxpool.Pool)

	poll := NewPoll{}
	err := c.ShouldBind(&poll)
	if err != nil {
		log.Println(err)
		return
	}

	query := `
	insert into polls_anketovac (
	title,
	description,
	ptype
	)
	values (
		$1,
		$2,
		$3
	)
	returning id
	;
	`
	var poll_id string
	err = dbpool.QueryRow(context.Background(), query, poll.Title, poll.Description, poll.Ptype).Scan(&poll_id)
	if err != nil {
		log.Println(err)
		return
	}
	query = `
	insert into poll_options (
	poll_id,
	name,
	option_order
	)
	values (
		$1,
		$2,
		$3
	)
	;
	`
	for i, option := range poll.Options {
		_, err = dbpool.Exec(context.Background(), query, poll_id, option, i)
		if err != nil {
			log.Println(err)
			return
		}
	}

	c.Redirect(304, "/")

}

func Delete(c *gin.Context) {}

func Edit(c *gin.Context) {}
