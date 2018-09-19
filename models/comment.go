package models

import (
	"time"

	"github.com/astaxie/beego/orm"
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

type CommentDBO struct{}

func (d *CommentDBO) AddComment(c Comment) error {
	O = orm.NewOrm()
	id, err := O.Insert(&c)
	if err != nil {
		return err
	}
	L.Printf("add comments success:%v\n", id)
	return nil
}

func (d *CommentDBO) GetCommentById(id int64) (Comment, error) {
	O = orm.NewOrm()
	c := Comment{Id: id, Deleted: 0}
	err := O.QueryTable(c.TableName()).Filter("id", id).Filter("deleted", 0).One(&c)
	if err != nil {
		L.Println("查询数据异常")
	}
	return c, err
}

func (d *CommentDBO) ListComments() ([]Comment, error) {
	O = orm.NewOrm()
	var cs []Comment
	_, err := O.QueryTable("t_comments").Filter("deleted", 0).All(&cs)
	return cs, err
}
