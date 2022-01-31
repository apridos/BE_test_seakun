package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" gorm:"size:20,unique"`
	Password string `gorm:"size:25"`

	TeacherId int
	Teacher   Teacher `gorm:"foreignKey:TeacherId"`
}

func (b *User) TableName() string {
	return "user"
}
