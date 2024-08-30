package article

type NewArticle struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required"`
	Summary     string `json:"summary" validate:"required"`
	Image       string `json:"-"`
	Category    uint   `json:"category" validate:"required,gt=0"`
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

type MinimumArtilce struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Image   string `json:"image"`
	ID      uint   `json:"id"`
}

type ArticleComment struct {
}
