package models

type Group struct {
	GroupID uint   `json:"id"`
	OwnerID uint   `json:"ownerId"`
	Name    string `json:"name"`
}
