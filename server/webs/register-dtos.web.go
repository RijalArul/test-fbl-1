package webs

type RegisterDTO struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserResponseBody struct {
	Username string `json:"username" form:"username"`
	Role     string `json:"role" form:"role"`
}

type LoginDTO struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResponseBody struct {
	AccessToken string `json:"access_token" form:"access_token"`
}
