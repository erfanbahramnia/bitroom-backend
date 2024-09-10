package developer

import (
	"bitroom/constants"
	user_model "bitroom/models/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DeveloperApi struct {
	db  *gorm.DB
	ech *echo.Echo
}

func NewDeveloperApi(db *gorm.DB, ech *echo.Echo) *DeveloperApi {
	return &DeveloperApi{
		db:  db,
		ech: ech,
	}
}

func (d *DeveloperApi) InitApi() {
	group := d.ech.Group("developer")

	group.PUT("/change-role", d.ChangeRole)
	group.GET("/users", d.GetUsers)
}

// @Summary change user role
// @Tags developer
// @Accept json
// @Produce json
// @Param register body ChangeRole true "Changing user role"
// @Router /developer/change-role [put]
func (d *DeveloperApi) ChangeRole(ctx echo.Context) error {
	var data ChangeRole

	// bind json to struct
	if err := ctx.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputFormat)
	}

	// update
	d.db.Model(&user_model.User{}).Where("id = ?", data.UserID).Updates(&user_model.User{Role: data.Role})

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "ok",
	})
}

// @Summary get all users
// @Tags developer
// @Produce json
// @Router /developer/users [get]
func (d *DeveloperApi) GetUsers(ctx echo.Context) error {
	var users []user_model.User

	if err := d.db.Model(&user_model.User{}).Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, constants.InternalServerError)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}
