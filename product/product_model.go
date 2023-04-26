package product

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type GetAllProductsRequest struct {
}

type GetAllProductsResponse struct {
	Products []Product `json:"products"`
}

type GetProductByIDRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetProductByIDResponse struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Product *Product `json:"product"`
}

type GetProductByNameRequest struct {
	Name string `json:"name" validate:"required"`
}

type GetProductByNameResponse struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Product *Product `json:"product"`
}
