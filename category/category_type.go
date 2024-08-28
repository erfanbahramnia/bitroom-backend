package category

type NewCategory struct {
	Name     string `json:"name" validate:"required,max=50"`
	ParentID *uint  `json:"parent_id" validate:"omitempty,gte=0"`
}

type EditCategory struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}
