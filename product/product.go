package product

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var products []Product

func SetUpProductsData() {
	for i := 1; i <= 10; i++ {
		id := uuid.New()
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		price := r1.Intn(100)

		product := Product{
			ID:    id.String(),
			Name:  fmt.Sprintf("Product_%d", i),
			Price: price,
		}

		products = append(products, product)

		fmt.Printf("ProductID: %s\n", product.ID)
	}
}

func GetAllProducts(c echo.Context) (err error) {

	req := new(GetAllProductsRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	resp := GetAllProductsResponse{
		Products: products,
	}

	return c.JSON(http.StatusOK, resp)
}

func GetProductByID(c echo.Context) (err error) {

	req := new(GetProductByIDRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var product *Product
	for _, p := range products {
		if req.ID == p.ID {
			product = &p
			break
		}
	}

	var resp GetProductByIDResponse

	if product == nil {
		resp = GetProductByIDResponse{
			Code:    "404",
			Message: "Product not found",
			Product: nil,
		}

		return c.JSON(http.StatusOK, resp)
	}

	resp = GetProductByIDResponse{
		Code:    "200",
		Message: "Success",
		Product: product,
	}

	return c.JSON(http.StatusOK, resp)
}

func GetProductByName(c echo.Context) (err error) {

	req := new(GetProductByNameRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var product *Product
	for _, p := range products {
		if req.Name == p.Name {
			product = &p
			break
		}
	}

	var resp GetProductByNameResponse

	if product == nil {
		resp = GetProductByNameResponse{
			Code:    "404",
			Message: "Product not found",
			Product: nil,
		}

		return c.JSON(http.StatusOK, resp)
	}

	resp = GetProductByNameResponse{
		Code:    "200",
		Message: "Success",
		Product: product,
	}

	return c.JSON(http.StatusOK, resp)
}
