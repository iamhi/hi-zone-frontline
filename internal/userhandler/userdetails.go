package userhandler

const USER_ROLE_BASIC = "Basic"
const USER_ROLE_SERVICE = "Service"

type UserDetails struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Roles    string `json:"roles"`
}
