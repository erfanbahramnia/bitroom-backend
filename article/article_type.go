package article

type NewArticle struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required"`
	Summary     string `json:"summary" validate:"required"`
	Image       string `json:"-"`
}

type UserComment struct {
}

type Article struct {
	Title       string
	Description string
	Summary     string
	Status      string
	Image       string
	ID          uint
}

type ArticleComment struct {
}
