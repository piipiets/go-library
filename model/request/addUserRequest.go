package request

type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
