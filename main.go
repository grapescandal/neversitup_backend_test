package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"neversitup_backend_test/auth"
	order "neversitup_backend_test/order"
	product "neversitup_backend_test/product"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func main() {
	productService := product.NewProductService()
	productService.SetUpProductsData()

	userService := order.NewUserService()
	userService.SetUpUsersProfile()

	orderService := order.NewOrderService(userService, productService)
	orderService.SetUpOrdersData()

	authService := auth.NewAuthService(userService)
	authService.SetUpAccountData()

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	r := e.Group("/user")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	//auth
	e.POST("/register", authService.Register)
	e.POST("/login", authService.Login)

	//product
	r.GET("/getAllProducts", productService.GetAllProducts)
	r.GET("/getProductByID", productService.GetProductByID)
	r.GET("/getProductByName", productService.GetProductByName)

	//user
	r.GET("/getUserProfileByID", userService.GetUserProfileByID)
	r.GET("/getUserOrderHistoryByID", userService.GetUserOrderHistoryByID)

	//order
	r.POST("/saveOrder", orderService.SaveOrder)
	r.GET("/getOrderByOrderID", orderService.GetOrderByOrderID)
	r.POST("/cancelOrderByOrderID", orderService.CancelOrderByOrderID)

	e.Logger.Fatal(e.Start(":8080"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
