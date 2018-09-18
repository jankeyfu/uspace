package models

import (
	"time"
)

type Comment struct {
	Id          int64 `orm:"auto"`
	Pid         int64
	Content     string
	FromUserId  int64
	ToUserId    int64
	CreatedTime time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedTime time.Time `orm:"auto_now_add;type(timestamp)"`
	Deleted     int8      `orm:"default(0)"`
}

func (c *Comment) TableName() string {
	return "t_comments"
}
