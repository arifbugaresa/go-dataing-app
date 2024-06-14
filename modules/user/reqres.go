package user

type UserReq struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Hobbies  string `json:"hobbies"`
	Password string `json:"password"`
}

type UserResp struct {
}

type LoginResp struct {
	Token string `json:"token"`
}
