package entity

type Person struct {
	ID        uint64 `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string `json:"firstname" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" gorm:"type:varchar(32)"`
	Age       int    `json:"age"`
	Email     string `json:"email" gorm:"type:varchar(256),unique"`
}
