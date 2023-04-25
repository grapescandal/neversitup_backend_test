package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	product "neversitup_backend_test/product"
	user "neversitup_backend_test/user"
)

func main() {

	product.SetUpProductsData()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	//product
	e.GET("/getAllProducts", product.GetAllProducts)
	e.GET("/getProductByID", product.GetProductByID)
	e.GET("/getProductByName", product.GetProductByName)

	//user
	e.GET("/getUserProfileByID", user.GetUserProfileByID)
	e.GET("/getUserOrderHistoryByID", user.GetUserOrderHistoryByID)
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
