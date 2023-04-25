package user

import (
	"fmt"
	"net/http"
	order "neversitup_backend_test/order"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var usersProfile []UserProfile

func SetUpUsersProfile() {
	usersProfile = []UserProfile{}
	fmt.Println("Set up UsersProfile completed")
}

func CreateUserProfile(name, lastName string) {
	id := uuid.New()
	usersProfile = append(usersProfile, UserProfile{
		ID:       id.String(),
		Name:     name,
		LastName: lastName,
	})
}

func GetUserProfileByID(c echo.Context) (err error) {

	req := new(GetUserProfileByIDRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var userProfile *UserProfile
	for _, u := range usersProfile {
		if req.ID == u.ID {
			userProfile = &u
			break
		}
	}

	var resp GetUserProfileByIDResponse

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

func GetUserOrderHistoryByID(c echo.Context) (err error) {

	req := new(GetUserOrderHistoryByIDRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var orders []order.Order
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
