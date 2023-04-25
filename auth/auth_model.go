package auth

import (
	user "neversitup_backend_test/user"
)

type RegisterRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

type RegisterResponse struct {
	Code        string            `json:"code"`
	Message     string            `json:"message"`
	UserName    string            `json:"userName"`
	UserProfile *user.UserProfile `json:"userProfile"`
}

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	JWT     string `json:"jwt"`
}
