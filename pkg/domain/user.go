package domain

type Users struct {
	Name     string `Json:"id"`
	Email    string `json:"email" gorm:"primary key"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Admin struct {
	Name     string
	Password string
}
