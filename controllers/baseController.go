package controllers

import (
	"uspace/models"
)

var (
	commentDBO models.CommentDBO
)

func init() {
	commentDBO = models.CommentDBO{}
}
