package user

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Hobbies  string `json:"hobbies"`
}
