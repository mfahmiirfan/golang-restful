package domain

type Gift struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Tag  string `gorm:"not null"`
	// Ratings     []Rating
	Points      int    `gorm:"not null"`
	Stock       int    `gorm:"not null"`
	Status      string `gorm:"not null"`
	Description string `gorm:"not null"`
	Image       string `gorm:"not null"`
}

func (p *Gift) TableName() string {
	return "gift"
}
