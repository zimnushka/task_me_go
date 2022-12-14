package usecases

import (
	"errors"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func (useCase *UserUseCase) GetUserById(id int) (*models.User, error) {
	return useCase.userRepository.GetUserFromId(id)
}

func (useCase *UserUseCase) GetUserByEmail(email string) (*models.User, error) {
	return useCase.userRepository.GetUserFromEmail(email)
}

func (useCase *UserUseCase) GetAllUsers() ([]models.User, error) {
	return useCase.userRepository.GetUsers()
}
func (useCase *UserUseCase) AddUser(user models.User) (*models.User, error) {
	user.Id = nil
	userWithEmail, _ := useCase.userRepository.GetUserFromEmail(user.Email)
	if userWithEmail != nil {
		return nil, errors.New("User with this email was created")
	}

	return useCase.userRepository.AddUser(user)
}
func (useCase *UserUseCase) UpdateUser(user models.User) (*models.User, error) {
	if user.Password == "" {
		userWithPass, _ := useCase.userRepository.GetUserFromId(*user.Id)
		user.Password = *&userWithPass.Password
	}

	err := useCase.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return useCase.userRepository.GetUserFromId(*user.Id)
}
func (useCase *UserUseCase) DeleteUser(id int) error {
	return useCase.userRepository.DeleteUser(id)
}
