package models

type PaymentMethod struct {
	PaymentMethodID uint   `json:"id"`
	UserID          uint   `json:"UserId"`
	Type            string `json:"type"`
	ProviderToken   string `json:"providerToken"`
}
