package order

import (
	product "neversitup_backend_test/product"
)

type Order struct {
	OrderID string           `json:"orderID"`
	Product *product.Product `json:"product"`
	UserID  string           `json:"userID"`
	Status  OrderStatus      `json:"status"`
}

type OrderStatus int64

const (
	Pending OrderStatus = iota
	Delivered
	Cancelled
)

type SaveOrderRequest struct {
	UserID    string `json:"userID"`
	ProductID string `json:"productID"`
}

type SaveOrderResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetOrderByOrderIDRequest struct {
	OrderID string `json:"orderID"`
}

type GetOrderByOrderIDResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Order   *Order `json:"order"`
}

type CancelOrderByOrderIDRequest struct {
	OrderID string `json:"orderID"`
}

type CancelOrderByOrderIDResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
