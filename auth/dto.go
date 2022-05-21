package auth

type SignupDto struct {
	Email                string `json:"email"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type SignupOutput struct {
	Token string `json:"token"`
}

func NewSignupOutput(token string) *SignupOutput {
	return &SignupOutput{token}
}
