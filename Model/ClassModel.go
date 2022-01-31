package model

type Class struct {
	Id   int    `json:"id"`
	Name string `json:"name" gorm:"size:5"`
}

func (b *Class) TableName() string {
	return "class"
}
