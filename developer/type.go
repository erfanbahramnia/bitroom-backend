package developer

type ChangeRole struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
}
