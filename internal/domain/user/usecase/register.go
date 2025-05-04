package usecase

import (
	"strings"

	"Go_API/internal/domain/user"
)

type RegisterUseCase struct {
	repo user.Repository
}

func NewRegisterUseCase(repo user.Repository) *RegisterUseCase {
	return &RegisterUseCase{repo: repo}
}

type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

type RegisterOutput struct {
	User *user.SafeUser
	Err  error
}

func (uc *RegisterUseCase) Execute(input RegisterInput) RegisterOutput {
	// Validation
	if strings.TrimSpace(input.Name) == "" {
		return RegisterOutput{Err: ErrNameRequired}
	}
	if strings.TrimSpace(input.Email) == "" {
		return RegisterOutput{Err: ErrEmailRequired}
	}
	if len(input.Password) < 6 {
		return RegisterOutput{Err: ErrPasswordTooShort}
	}

	// Create user
	user := &user.User{
		Name:     strings.TrimSpace(input.Name),
		Email:    strings.TrimSpace(input.Email),
		Password: input.Password,
	}

	// Save to repository
	if err := uc.repo.Create(user); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return RegisterOutput{Err: ErrEmailExists}
		}
		return RegisterOutput{Err: err}
	}

	return RegisterOutput{User: user.ToSafeUser()}
}
