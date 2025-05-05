package usecase

import user "github.com/aleksander/Go_API/internal/domain/user/models"

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
	User *user.SafeUser
	Err  error
}

func (uc *GetUserUseCase) Execute(input GetUserInput) GetUserOutput {
	user, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return GetUserOutput{Err: ErrUserNotFound}
	}
	return GetUserOutput{User: user.ToSafeUser()}
}

type ListUsersUseCase struct {
	repo user.Repository
}

func NewListUsersUseCase(repo user.Repository) *ListUsersUseCase {
	return &ListUsersUseCase{repo: repo}
}

type ListUsersOutput struct {
	Users []*user.SafeUser
	Err   error
}

func (uc *ListUsersUseCase) Execute() ListUsersOutput {
	users, err := uc.repo.FindAll()
	if err != nil {
		return ListUsersOutput{Err: err}
	}

	safeUsers := make([]*user.SafeUser, len(users))
	for i, user := range users {
		safeUsers[i] = user.ToSafeUser()
	}
	return ListUsersOutput{Users: safeUsers}
}
