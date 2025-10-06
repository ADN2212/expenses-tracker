package IO

type TransactionInput struct {
	Description string `json:"description" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
}
