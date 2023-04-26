package auth

import (
	"fmt"
	"net/http"
	"neversitup_backend_test/order"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var accounts []Account

type AuthService struct {
	userService *order.UserService
}

func NewAuthService(userService *order.UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

func (a AuthService) SetUpAccountData() {
	accounts = []Account{}
	fmt.Println("Set up AccountData completed")
}

func (a AuthService) GetAccountDataByUserName(username string) *Account {
	for _, a := range accounts {
		if username == a.Username {
			return &a
		}
	}

	return nil
}

func (a AuthService) RegisterAccountData(username, password string) (userID string) {
	id := uuid.New()
	account := Account{
		Username: username,
		Password: password,
		UserID:   id.String(),
	}
	accounts = append(accounts, account)

	return account.UserID
}

func (a AuthService) Register(c echo.Context) (err error) {

	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	userID := a.RegisterAccountData(req.UserName, req.Password)
	newUserProfile := a.userService.CreateUserProfile(userID, req.Name, req.Lastname)

	resp := RegisterResponse{
		Code:        "200",
		Message:     "Success",
		UserName:    req.UserName,
		UserProfile: &newUserProfile,
	}

	return c.JSON(http.StatusOK, resp)
}

func (a AuthService) Login(c echo.Context) (err error) {

	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var resp LoginResponse

	account := a.GetAccountDataByUserName(req.Username)

	if account == nil {
		resp = LoginResponse{
			Code:    "404",
			Message: "Username not found",
		}

		return c.JSON(http.StatusOK, resp)
	}

	tokenWithSecret := ""
	if req.Password != account.Password {
		resp = LoginResponse{
			Code:    "404",
			Message: "Username or password not correct",
		}

		return c.JSON(http.StatusOK, resp)
	} else {
		userProfile := a.userService.GetUserProfileByIDInternal(account.UserID)
		// Set custom claims
		claims := &jwtCustomClaims{
			Name:  userProfile.Name,
			Admin: false,
			JWT: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.JWT)

		// Generate encoded token and send it as response.
		tokenWithSecret, err = token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
	}

	resp = LoginResponse{
		Code:    "200",
		Message: "Success",
		Token:   tokenWithSecret,
	}

	return c.JSON(http.StatusOK, resp)
}
