package auth

import (
	"bitroom/constants"
	"bitroom/types"
	"bitroom/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service AuthServiceInterface
}

func NewAuthHandler(service AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (a *AuthHandler) InitHandler(ech *echo.Echo) {
	group := ech.Group("auth")

	group.POST("/login/password", a.LoginWithPassword)
	group.POST("/login/send-otp", a.SendOtpForLoging)
	group.POST("/login/validate-otp", a.OtpLoginValidation)
	group.POST("/register/send-otp", a.OtpRegistering)
	group.POST("/register/validate-otp", a.OtpRegisterValidation)
}

// @Summary Register
// @Description User registration
// @Tags auth
// @Accept json
// @Produce json
// @Param register body RequiredDataForOtp true "Register request"
// @Success 201
// @Router /auth/register/send-otp [post]
func (a *AuthHandler) OtpRegistering(ctx echo.Context) error {
	var userData RequiredDataForOtp

	// bind json to struct
	if err := ctx.Bind(&userData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputFormat)
	}

	// validate
	vs := utils.GetValidator()
	vsErrs := vs.Validate(userData)
	if len(vsErrs) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"errors": vsErrs,
		})
	}

	// get otp
	otp, err := a.service.OtpGeneratingForRegister(userData.Phone)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	// send otp via sms
	fmt.Println("phone", userData.Phone, "otp", otp)

	// success
	return ctx.JSON(http.StatusOK, constants.OtpSended)
}

// @Summary Register
// @Description User otp validation
// @Tags auth
// @Accept json
// @Produce json
// @Param register body ValidateOtp true "Register request"
// @Success 201 {object} RegisterResponse
// @Router /auth/register/validate-otp [post]
func (a *AuthHandler) OtpRegisterValidation(ctx echo.Context) error {
	var data ValidateOtp

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

	// create new user
	user, err := a.service.ValidateOtpRegister(data)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}
	// generate new jwt tokens
	claims := types.UserDataJwtClaims{
		Id:    user.ID,
		Role:  user.Role,
		Phone: data.Phone,
	}
	jwt, jwtErr := utils.GenerateJwt(claims)
	if jwtErr != nil {
		fmt.Println(jwtErr)
		return echo.NewHTTPError(http.StatusInternalServerError, constants.InternalServerError)
	}

	res := &AuthResponse{
		Phone:      user.Phone,
		First_name: user.FirstName,
		Last_name:  user.LastName,
		ID:         user.ID,
		Role:       user.Role,
		Jwt: types.JwtTokens{
			Token:        jwt.Token,
			RefreshToken: jwt.RefreshToken,
		},
	}
	return ctx.JSON(http.StatusOK, res)
}

// @Summary login with otp
// @Tags auth
// @Accept json
// @Produce json
// @Param register body RequiredDataForOtp true "Register request"
// @Router /auth/login/send-otp [post]
func (a *AuthHandler) SendOtpForLoging(ctx echo.Context) error {
	var userData RequiredDataForOtp

	// bind json to struct
	if err := ctx.Bind(&userData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputFormat)
	}

	// validate
	vs := utils.GetValidator()
	vsErrs := vs.Validate(userData)
	if len(vsErrs) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"errors": vsErrs,
		})
	}

	//

	// get otp
	otp, err := a.service.OtpGeneratingForLogin(userData.Phone)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	// send otp via sms
	fmt.Println("phone", userData.Phone, "otp", otp)

	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": constants.OtpSended,
	})
}

// @Description send otp for loging
// @Tags auth
// @Accept json
// @Produce json
// @Param register body ValidateOtp true "Login request"
// @Router /auth/login/validate-otp [post]
func (a *AuthHandler) OtpLoginValidation(ctx echo.Context) error {
	var data ValidateOtp

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

	// create new user
	user, err := a.service.ValidateOtpLogin(data)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}
	// generate new jwt tokens
	claims := types.UserDataJwtClaims{
		Id:    user.ID,
		Role:  user.Role,
		Phone: data.Phone,
	}
	jwt, jwtErr := utils.GenerateJwt(claims)
	if jwtErr != nil {
		fmt.Println(jwtErr)
		return echo.NewHTTPError(http.StatusInternalServerError, constants.InternalServerError)
	}

	res := &AuthResponse{
		Phone:      user.Phone,
		First_name: user.FirstName,
		Last_name:  user.LastName,
		ID:         user.ID,
		Role:       user.Role,
		Jwt: types.JwtTokens{
			Token:        jwt.Token,
			RefreshToken: jwt.RefreshToken,
		},
	}
	return ctx.JSON(http.StatusOK, res)
}

// @Description login with password
// @Tags auth
// @Accept json
// @Produce json
// @Param register body LoginCredential true "login with password"
// @Router /auth/login/password [post]
func (a *AuthHandler) LoginWithPassword(ctx echo.Context) error {
	var data LoginCredential

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
	user, err := a.service.LoginWithPassword(&data)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	claim := types.UserDataJwtClaims{
		Id:    user.ID,
		Role:  user.Role,
		Phone: user.Phone,
	}
	jwt, jwtErr := utils.GenerateJwt(claim)
	if jwtErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, constants.InternalServerError)
	}
	res := &AuthResponse{
		Phone:      user.Phone,
		First_name: user.FirstName,
		Last_name:  user.LastName,
		ID:         user.ID,
		Role:       user.Role,
		Jwt: types.JwtTokens{
			Token:        jwt.Token,
			RefreshToken: jwt.RefreshToken,
		},
	}
	return ctx.JSON(http.StatusOK, res)
}
