package models

type Subscription struct {
	SubscriptionID uint   `json:"id"`
	GroupID        uint   `json:"groupId"`
	ServiceID      uint   `json:"serviceId"`
	StartDate      string `json:"startDate"`
	CycleInterval  string `json:"cycleInterval"`
	NextDueDate    string `json:"dueDate"`
}
