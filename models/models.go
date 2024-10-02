package models

import "time"

type PType string

type Poll struct {
	Id          string
	Title       string
	Description string
	Ptype       PType
	Created_at  time.Time
}

const (
	Multiple PType = "multiple"
	Ranking  PType = "ranking"
	Image    PType = "image"
)

type PollOption struct {
	Poll_id     string
	OptionOrder uint
	Name        string
	Votes       uint
}

type NewPoll struct {
	Title       string   `form:"title" json:"title" xml:"title"  binding:"required"`
	Description string   `form:"description" json:"description" xml:"description"  binding:"required"`
	Ptype       PType    `form:"ptype" json:"ptype" xml:"ptype"  binding:"required"`
	Options     []string `form:"options" json:"options" xml:"options"  binding:"required"`
}
