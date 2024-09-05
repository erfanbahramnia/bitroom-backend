package user

import (
	"bitroom/constants"
	"bitroom/middleware"
	"bitroom/types"
	"bitroom/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service UserServiceInterface
}

func NewUserHandler(service UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (u *UserHandler) InitHandler(ech *echo.Echo) {
	group := ech.Group("user", middleware.JwtMiddleware)

	group.PUT("/edit", u.EditUser)
	group.PUT("/password/change", u.ChangePaasword)
}

// --------------------------------------------------------------------------------------------------------------------

// @Description edit user data
// @Tags users
// @Accept json
// @Produce json
// @Param register body EditUser true "Edit data"
// @Router /user/edit [put]
// @Security BearerAuth
func (u *UserHandler) EditUser(ctx echo.Context) error {
	var data EditUser

	// bind json to struct
	if err := ctx.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputFormat)
	}

	// validate
	vs := utils.GetValidator()
	vsErrs := vs.Validate(data)
	if len(vsErrs) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"errors": vsErrs,
		})
	}

	// get user id
	claim := ctx.Get("user").(*types.JwtCustomClaims)
	userid := claim.Id

	// update
	err := u.service.EditUserData(&data, userid)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	return ctx.JSON(http.StatusOK, utils.MapStringInterface("msg", "ok"))
}

// --------------------------------------------------------------------------------------------------------------------

// @Description change password
// @Tags users
// @Accept json
// @Produce json
// @Param register body ChangePaasword true "Change password"
// @Router /user/password/change [put]
// @Security BearerAuth
func (u *UserHandler) ChangePaasword(ctx echo.Context) error {
	var data ChangePaasword

	// bind json to struct
	if err := ctx.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputFormat)
	}

	// validate
	vs := utils.GetValidator()
	vsErrs := vs.Validate(data)
	if len(vsErrs) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"errors": vsErrs,
		})
	}

	// get user data
	claim := ctx.Get("user").(*types.JwtCustomClaims)

	// update
	if err := u.service.ChangePaasword(claim.Phone, data.Password); err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "password updated successfully",
	})
}
