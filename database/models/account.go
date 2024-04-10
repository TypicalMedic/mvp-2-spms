package models

type Account struct {
	Id    uint   `gorm:"column:professor_id"`
	Login string `gorm:"column:login"`
	Hash  string `gorm:"column:hash"`
}

func (Account) TableName() string {
	return "user_account"
}
