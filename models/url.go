package models

// URLMapping is the model for storing URL mappings.
type URLMapping struct {
	ShortCode string `gorm:"primaryKey;size:6"`
	LongURL   string `gorm:"not null"`
}
