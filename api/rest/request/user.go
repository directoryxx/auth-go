package request

type UserRequest struct {
	Name     string
	Username string
	Password string
	RoleId   int
}

type RegisterUserRequest struct {
	Name     string
	Username string
	Password string
}

type LoginUserRequest struct {
	Username string
	Password string
}
