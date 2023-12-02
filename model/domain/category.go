package domain

type Category struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (p *Category) TableName() string {
	return "category"
}
