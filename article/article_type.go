package article

type NewArticle struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required"`
	Summary     string `json:"summary" validate:"required"`
	Image       string `json:"-"`
	Category    uint   `json:"category" validate:"required,gt=0"`
}

type ArticleProperty struct {
	ArticleID   uint    `json:"article_id" form:"article_id" validate:"required"`
	Description string  `json:"description" form:"description" validate:"required,min=10"`
	Image       *string `json:"image" form:"image"`
}

type UserComment struct {
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	Status      string `json:"status"`
	Image       string `json:"image"`
	ID          uint   `json:"id"`
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

type ArticleComment struct {
}
