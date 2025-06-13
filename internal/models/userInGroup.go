package models

type UserInGroup struct {
	UserInGroupID uint   `json:"id"`
	GroupID       uint   `json:"groupId"`
	UserID        uint   `json:"userId"`
	Role          string `json:"role"`
	IsPaying      bool   `json:"isPaying"`
	FlexPaid      bool   `json:"flexPaid"`
	Status        string `json:"status"`
	Percent       uint   `json:"percent"`
	Amount        uint   `json:"amount"`
}
