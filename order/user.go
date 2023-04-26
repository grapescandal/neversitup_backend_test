package order

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var usersProfile []UserProfile

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) SetUpUsersProfile() {
	usersProfile = []UserProfile{
		{
			ID:       "1",
			Name:     "Grape",
			Lastname: "Test",
		},
	}
	fmt.Println("Set up UsersProfile completed")
}

func (u UserService) CreateUserProfile(userID, name, lastName string) UserProfile {
	userProfle := UserProfile{
		ID:       userID,
		Name:     name,
		Lastname: lastName,
	}
	usersProfile = append(usersProfile, userProfle)

	return userProfle
}

func (u UserService) GetUserProfileByIDInternal(id string) *UserProfile {
	for _, u := range usersProfile {
		if id == u.ID {
			return &u
		}
	}

	return nil
}

func (u UserService) GetUserProfileByID(c echo.Context) (err error) {

	req := new(GetUserProfileByIDRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var resp GetUserProfileByIDResponse

	userProfile := u.GetUserProfileByIDInternal(req.ID)

	if userProfile == nil {
		resp = GetUserProfileByIDResponse{
			Code:        "404",
			Message:     "UserProfile not found",
			UserProfile: nil,
		}

		return c.JSON(http.StatusOK, resp)
	}

	resp = GetUserProfileByIDResponse{
		Code:        "200",
		Message:     "Success",
		UserProfile: userProfile,
	}

	return c.JSON(http.StatusOK, resp)
}

func (u UserService) GetUserOrderHistoryByID(c echo.Context) (err error) {

	req := new(GetUserOrderHistoryByIDRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var orders []Order
	var resp GetUserOrderHistoryByIDResponse

	if orders == nil {
		resp = GetUserOrderHistoryByIDResponse{
			Code:    "404",
			Message: "UserProfile not found",
			Orders:  nil,
		}

		return c.JSON(http.StatusOK, resp)
	}

	resp = GetUserOrderHistoryByIDResponse{
		Code:    "200",
		Message: "Success",
		Orders:  orders,
	}

	return c.JSON(http.StatusOK, resp)
}
