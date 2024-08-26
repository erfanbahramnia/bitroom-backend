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

	group.POST("/login", a.HandleLogin)
	group.POST("/register/send-otp", a.OtpRegistering)
	group.POST("/register/validate-otp", a.OtpRegisterValidation)
}

// @Summary Register
// @Description User registration
// @Tags auth
// @Accept json
// @Produce json
// @Param register body RequiredOtpRegistering true "Register request"
// @Success 201
// @Router /auth/register/send-otp [post]
func (a *AuthHandler) OtpRegistering(ctx echo.Context) error {
	var userData RequiredOtpRegistering

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
// @Param register body ValidateOtpRegistering true "Register request"
// @Success 201 {object} RegisterResponse
// @Router /auth/register/validate-otp [post]
func (a *AuthHandler) OtpRegisterValidation(ctx echo.Context) error {
	var data ValidateOtpRegistering

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

// @Summary Login
// @Description User registration
// @Tags auth
// @Accept json
// @Produce json
// @Param register body LoginCredential true "Register request"
// @Success 201 {object} AuthResponse
// @Router /auth/login [post]
func (a *AuthHandler) HandleLogin(ctx echo.Context) error {
	var userReq LoginCredential

	// bind json to struct
	if err := ctx.Bind(&userReq); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputFormat)
	}

	// validate data
	vs := utils.GetValidator()
	vsErrs := vs.Validate(userReq)
	if len(vsErrs) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"errors": vsErrs,
		})
	}

	// get user data
	userData, err := a.service.Login(userReq)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}
	// generate new jwt tokens
	claims := types.UserDataJwtClaims{
		Phone: userReq.Phone,
	}
	jwt, jwtErr := utils.GenerateJwt(claims)
	if jwtErr != nil {
		fmt.Println(jwtErr)
		return echo.NewHTTPError(http.StatusInternalServerError, constants.InternalServerError)
	}

	res := &AuthResponse{
		Phone:      userData.Phone,
		First_name: userData.FirstName,
		Last_name:  userData.LastName,
		ID:         userData.ID,
		Role:       userData.Role,
		Jwt: types.JwtTokens{
			Token:        jwt.Token,
			RefreshToken: jwt.RefreshToken,
		},
	}
	return ctx.JSON(http.StatusOK, res)
}
