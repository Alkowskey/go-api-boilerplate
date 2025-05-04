package usecase

import (
	"Go_API/internal/domain/user"
)

type GetUserUseCase struct {
	repo user.Repository
}

func NewGetUserUseCase(repo user.Repository) *GetUserUseCase {
	return &GetUserUseCase{repo: repo}
}

type GetUserInput struct {
	ID uint
}

type GetUserOutput struct {
	User *user.User
	Err  error
}

func (uc *GetUserUseCase) Execute(input GetUserInput) GetUserOutput {
	user, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return GetUserOutput{Err: ErrUserNotFound}
	}
	user.Password = ""
	return GetUserOutput{User: user}
}

type ListUsersUseCase struct {
	repo user.Repository
}

func NewListUsersUseCase(repo user.Repository) *ListUsersUseCase {
	return &ListUsersUseCase{repo: repo}
}

type ListUsersOutput struct {
	Users []user.User
	Err   error
}

func (uc *ListUsersUseCase) Execute() ListUsersOutput {
	users, err := uc.repo.FindAll()
	if err != nil {
		return ListUsersOutput{Err: err}
	}
	// Clear passwords
	for i := range users {
		users[i].Password = ""
	}
	return ListUsersOutput{Users: users}
}
