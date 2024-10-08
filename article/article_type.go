package article

import "github.com/lib/pq"

type NewArticle struct {
	Title       string `form:"title" validate:"required,max=100"`
	Description string `form:"description" validate:"required"`
	Summary     string `form:"summary" validate:"required"`
	Image       string `form:"-"`
	Category    uint   `form:"category" validate:"required,gt=0"`
}

type ArticleProperty struct {
	ArticleID   uint    `json:"article_id" form:"article_id" validate:"required"`
	Description string  `json:"description" form:"description" validate:"required,min=10"`
	Image       *string `json:"image" form:"image"`
}

type EditArticleProperty struct {
	PropertyID  uint    `json:"property_id" form:"property_id"`
	Description *string `json:"description" form:"description"`
	Image       *string `json:"image" form:"image"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	Status      string `json:"status"`
	Image       string `json:"image"`
	Likes       pq.Int64Array
	Dislikes    pq.Int64Array
	Views       int
	ID          uint `json:"id"`
}

type EditArticle struct {
	Title       *string `json:"title" form:"title"`
	Description *string `json:"description" form:"description"`
	Summary     *string `json:"summary" form:"summary"`
	Status      *string `json:"status" form:"status"`
	Image       *string `json:"image" form:"image"`
	Category    *uint   `json:"category" form:"category"`
	Id          *uint   `json:"id" form:"id"`
}

type MinimumArticle struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Image   string `json:"image"`
	ID      uint   `json:"id"`
}

type NewComment struct {
	Comment   string `json:"comment" validate:"required,min=3"`
	UserID    uint   `json:"-"`
	ArticleID uint   `json:"article_id" validate:"required"`
}

type EditComment struct {
	ArticleId uint   `json:"article_id" validate:"required"`
	CommentId uint   `json:"comment_id" validate:"required"`
	UserID    uint   `json:"-"`
	Comment   string `json:"comment" validate:"required,min=3"`
}

type EditCommentByAdmin struct {
	CommentId uint   `json:"comment_id" validate:"required"`
	Comment   string `json:"comment" validate:"required,min=3"`
}

type DeleteComment struct {
	ArticleId uint `json:"article_id" validate:"required"`
	CommentId uint `json:"comment_id" validate:"required"`
	UserID    uint `json:"-"`
}
