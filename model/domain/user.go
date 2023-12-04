package domain

import "github.com/google/uuid"

type User struct {
	ID       *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string     `gorm:"type:varchar(100);not null"`
	Email    string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	// Role      *string    `gorm:"type:varchar(50);default:'user';not null"`
	// Provider  *string    `gorm:"type:varchar(50);default:'local';not null"`
	// Photo     *string    `gorm:"not null;default:'default.png'"`
	// Verified  *bool      `gorm:"not null;default:false"`
	// CreatedAt *time.Time `gorm:"not null;default:now()"`
	// UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

func (p *User) TableName() string {
	return "user"
}
