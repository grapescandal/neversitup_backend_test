package user

import (
	order "neversitup_backend_test/order"
)

type UserProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

type GetUserProfileByIDRequest struct {
	ID string `json:"id"`
}

type GetUserProfileByIDResponse struct {
	Code        string       `json:"code"`
	Message     string       `json:"message"`
	UserProfile *UserProfile `json:"userProfile"`
}

type GetUserOrderHistoryByIDRequest struct {
	ID string `json:"id"`
}

type GetUserOrderHistoryByIDResponse struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Orders  []order.Order `json:"orders"`
}
