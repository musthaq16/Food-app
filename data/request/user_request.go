package request

type CreateUserRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required,min=3,max=100" json:"email"`
	Password string `validate:"required,min=4,max=100" json:"password"`
}

type UpdateUserRequest struct {
	Id       int    `validate:"required"`
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required,min=3,max=100" json:"email"`
	Password string `validate:"required,min=4,max=100" json:"password"`
}

type LoginRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	// Email    string `validate:"required,min=3,max=100" json:"email"`
	Password string `validate:"required,min=4,max=100" json:"password"`
}

type ForgetPasswordRequest struct {
	Email string `validate:"required,min=3,max=100" json:"email"`
}

type ResetPasswordRequest struct {
	Password        string `validate:"required,min=4,max=100" json:"password"`
	ConfirmPassword string `validate:"required,min=4,max=100" json:"confirmpassword"`
}
