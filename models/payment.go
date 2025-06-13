package models

type Payment struct {
	PaymentID      uint   `json:"id"`
	SubscriptionID uint   `json:"subscriptionId"`
	UserID         uint   `json:"userID"`
	Amount         uint   `json:"amount"`
	DueDate        string `json:"dueDate"`
	PaidAt         string `json:"paidAt"`
	Status         string `json:"status"`
}
