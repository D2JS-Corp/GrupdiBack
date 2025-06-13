package models

type Service struct {
	ServiceID    uint   `json:"id"`
	Name         string `json:"name"`
	BillingCycle string `json:"billing"`
}
