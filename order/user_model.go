package order

type UserProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type GetUserProfileByIDRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetUserProfileByIDResponse struct {
	Code        string       `json:"code"`
	Message     string       `json:"message"`
	UserProfile *UserProfile `json:"userProfile"`
}

type GetUserOrderHistoryByIDRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetUserOrderHistoryByIDResponse struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Orders  []Order `json:"orders"`
}
