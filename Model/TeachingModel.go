package model

type Teaching struct {
	TeacherId int `json:"teacher_id" gorm:"primaryKey"`
	ClassId   int `json:"class_id" gorm:"primaryKey"`

	Teacher Teacher `gorm:"foreignKey:TeacherId"`
	Class   Class   `gorm:"foreignKey:ClassId"`
}

func (b *Teaching) TableName() string {
	return "teaching"
}
