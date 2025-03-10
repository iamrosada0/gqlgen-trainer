package model

type Event struct {
	ID           string         `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Price        float64        `json:"price"`
	Date         *string        `json:"date,omitempty"`
	ImageURL     string         `json:"imageUrl"`
	LocationID   string         `gorm:"index" json:"locationId"`
	StreetImages []*StreetImage `gorm:"foreignKey:EventID" json:"streetImages"`
}

type StreetImage struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	EventID string `gorm:"index"`
	Event   *Event `gorm:"foreignKey:EventID"`
}
