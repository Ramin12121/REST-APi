package Server

type Subscription struct {
	ID          string `gorm:"primaryKey" json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date" example:"2025-09-01T00:00:00Z"`
	EndDate     string `json:"end_date,omitempty"`
}

type SubscriptionRequest struct {
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
}

type ToFilter struct {
	ServiceName string `json:"service_name"`
	UserID      string `json:"user_id"`
}
