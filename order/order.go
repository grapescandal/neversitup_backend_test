package order

import (
	"fmt"
	"net/http"
	"neversitup_backend_test/product"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var orders []Order

type OrderService struct {
	userService    *UserService
	productService *product.ProductService
}

func NewOrderService(userService *UserService, productService *product.ProductService) *OrderService {
	return &OrderService{
		userService:    userService,
		productService: productService,
	}
}

func (o OrderService) SetUpOrdersData() {
	orders = []Order{}
	fmt.Println("Set up OrdersData completed")
}

func (o OrderService) SaveOrderInternal(userID, productID string) string {
	id := uuid.New()
	order := Order{
		OrderID: id.String(),
		Product: productID,
		UserID:  userID,
	}

	order.Status = order.GetStatusString()
	orders = append(orders, order)

	return order.OrderID
}

func (o OrderService) SaveOrder(c echo.Context) (err error) {

	req := new(SaveOrderRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var resp SaveOrderResponse

	//Validate User
	userProfile := o.userService.GetUserProfileByIDInternal(req.UserID)

	if userProfile == nil {
		resp = SaveOrderResponse{
			Code:    "404",
			Message: "User not found",
		}

		return c.JSON(http.StatusOK, resp)
	}

	product := o.productService.GetProductByIDInternal(req.ProductID)

	if product == nil {
		resp = SaveOrderResponse{
			Code:    "404",
			Message: "Product not found",
			OrderID: "",
		}

		return c.JSON(http.StatusOK, resp)
	}

	orderID := o.SaveOrderInternal(req.UserID, req.ProductID)

	resp = SaveOrderResponse{
		Code:    "200",
		Message: "Success",
		OrderID: orderID,
	}

	return c.JSON(http.StatusOK, resp)
}

func (o OrderService) GetOrderByOrderIDInternal(id string) *Order {
	for i := range orders {
		if id == orders[i].OrderID {
			return &orders[i]
		}
	}

	return nil
}

func (o OrderService) GetOrderByOrderID(c echo.Context) (err error) {

	req := new(GetOrderByOrderIDRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var resp GetOrderByOrderIDResponse

	order := o.GetOrderByOrderIDInternal(req.OrderID)

	if order == nil {
		resp = GetOrderByOrderIDResponse{
			Code:    "404",
			Message: "Order not found",
			Order:   nil,
		}

		return c.JSON(http.StatusOK, resp)
	}

	resp = GetOrderByOrderIDResponse{
		Code:    "200",
		Message: "Success",
		Order:   order,
	}

	return c.JSON(http.StatusOK, resp)
}

func (o OrderService) CancelOrderInternal(order *Order) *Order {
	order.orderStatus = Cancelled
	order.Status = order.GetStatusString()
	return order
}

func (o OrderService) CancelOrderByOrderID(c echo.Context) (err error) {

	req := new(CancelOrderByOrderIDRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	var resp CancelOrderByOrderIDResponse

	order := o.GetOrderByOrderIDInternal(req.OrderID)

	if order == nil {
		resp = CancelOrderByOrderIDResponse{
			Code:    "404",
			Message: "Order not found",
		}

		return c.JSON(http.StatusOK, resp)
	}

	_ = o.CancelOrderInternal(order)
	fmt.Printf("Orders: %v\n", orders)

	resp = CancelOrderByOrderIDResponse{
		Code:    "200",
		Message: "Success",
	}

	return c.JSON(http.StatusOK, resp)
}
