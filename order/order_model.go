package order

type Order struct {
	OrderID     string `json:"orderID"`
	Product     string `json:"product"`
	UserID      string `json:"userID"`
	Status      string `json:"status"`
	orderStatus OrderStatus
}

type OrderStatus int64

const (
	Pending OrderStatus = iota
	Delivered
	Cancelled
)

func (o Order) GetStatusString() string {
	switch o.orderStatus {
	case Pending:
		return "Pending"
	case Delivered:
		return "Delivered"
	case Cancelled:
		return "Cancelled"
	default:
		return ""
	}
}

type SaveOrderRequest struct {
	UserID    string `json:"userID" validate:"required"`
	ProductID string `json:"productID" validate:"required"`
}

type SaveOrderResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	OrderID string `json:"orderID"`
}

type GetOrderByOrderIDRequest struct {
	OrderID string `json:"orderID" validate:"required"`
}

type GetOrderByOrderIDResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Order   *Order `json:"order"`
}

type CancelOrderByOrderIDRequest struct {
	OrderID string `json:"orderID" validate:"required"`
}

type CancelOrderByOrderIDResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
