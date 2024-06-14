package user

type UserReq struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Hobbies  string `json:"hobbies"`
	Password string `json:"password"`
}
