package services

import (
	"github.com/google/uuid"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GetUser(filter models.User) (*models.User, error) {
	return s.db.GetUser(&filter)
}

func (s *Service) LoginUser(username, password string) (*models.User, error) {
	user, err := s.db.LoginUser(username, password)
	if err != nil {
		log().Errorf("error getting user with given username and password: %v", err)
		return nil, err
	}
	user.Token = uuid.New().String() + "-" + uuid.New().String()
	err = s.db.UpdateUser(user)
	if err != nil {
		log().Errorf("error updating the user token: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *Service) UpdatePass(user *models.User, pass string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	err = s.db.UpdateUser(user)
	return user, err
}

func (s *Service) RegisterUser(user models.User) error {
	return s.db.Register(&user)
}
