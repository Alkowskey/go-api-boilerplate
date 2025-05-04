package usecase

import (
	"Go_API/internal/domain/user"
)

type AuthenticateUseCase struct {
	repo user.Repository
}

func NewAuthenticateUseCase(repo user.Repository) *AuthenticateUseCase {
	return &AuthenticateUseCase{repo: repo}
}

type AuthenticateInput struct {
	Email    string
	Password string
}

type AuthenticateOutput struct {
	User *user.User
	Err  error
}

func (uc *AuthenticateUseCase) Execute(input AuthenticateInput) AuthenticateOutput {
	// Find user by email
	user, err := uc.repo.FindByEmail(input.Email)
	if err != nil {
		return AuthenticateOutput{Err: ErrInvalidCredentials}
	}

	// Check password
	if !user.CheckPassword(input.Password) {
		return AuthenticateOutput{Err: ErrInvalidCredentials}
	}

	// Clear password before returning
	user.Password = ""
	return AuthenticateOutput{User: user}
}
