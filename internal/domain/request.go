package domain

type ReqRegisterUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type ReqLogin struct {
	Username string `json:"username" `
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" `
}

type ReqProduct struct {
	NameProduct string `json:"nameProduct" binding:"required"`
	Desciption  string
	Price       float64
	Stock       int
}

type ReqTransaction struct {
	ProductId uint `json:"productId"  binding:"required"`
	Quantity  int  `json:"quantity"  binding:"required"`
}
