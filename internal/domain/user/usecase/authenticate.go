package usecase

import user "github.com/aleksander/Go_API/internal/domain/user/models"

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
	User *user.SafeUser
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

	return AuthenticateOutput{User: user.ToSafeUser()}
}
