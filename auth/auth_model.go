package auth

import (
	user "neversitup_backend_test/order"

	"github.com/golang-jwt/jwt/v4"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserID   string `json:"userID"`
}

type RegisterRequest struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
}

type RegisterResponse struct {
	Code        string            `json:"code"`
	Message     string            `json:"message"`
	UserName    string            `json:"userName"`
	UserProfile *user.UserProfile `json:"userProfile"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	JWT   jwt.RegisteredClaims
}
