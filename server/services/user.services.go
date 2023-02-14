package services

import (
	"test-fbl-1/server/entities"
	"test-fbl-1/server/repositories"
	"test-fbl-1/server/webs"
)

func UserBodyResponse(user *entities.User) webs.UserResponseBody {
	return webs.UserResponseBody{
		Username: user.Username,
		Role:     user.Role,
	}
}

type UserService interface {
	Register(userDTO webs.RegisterDTO) (webs.UserResponseBody, error)
	Login(userDTO webs.LoginDTO) (*entities.User, error)
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(UserRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepository: UserRepository}
}

func (s *UserServiceImpl) Register(userDTO webs.RegisterDTO) (webs.UserResponseBody, error) {
	user := entities.User{
		Username: userDTO.Username,
		Password: userDTO.Password,
		Role:     "Admin",
	}
	createUser, err := s.userRepository.Create(user)
	userResp := UserBodyResponse(createUser)
	return userResp, err
}

func (s *UserServiceImpl) Login(userDTO webs.LoginDTO) (*entities.User, error) {
	username, err := s.userRepository.FindByUsername(userDTO.Username)
	return username, err
}
