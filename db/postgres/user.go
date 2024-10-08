package postgres

import (
	"errors"
	"fmt"

	wraperrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/hamza-sharif/homeopathic-doctor-assistant/config"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

func (cli *client) GetUser(userFilter *models.User) (*models.User, error) {
	var user *models.User
	result := cli.conn.Where(userFilter).Find(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("No user found with the specified username!")
			return nil, wraperrors.Wrap(result.Error, config.ErrorMessage404)
		} else {
			return nil, wraperrors.Wrap(result.Error, config.ErrorMessage500)
		}
	}

	return user, nil
}

func (cli *client) LoginUser(username, password string) (*models.User, error) {
	var user *models.User
	result := cli.conn.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("No user found with the specified username!")
			return nil, wraperrors.Wrap(result.Error, config.ErrorMessage404)
		} else {
			return nil, wraperrors.Wrap(result.Error, config.ErrorMessage500)
		}
	}
	// Compare the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Incorrect password!", err)
		return nil, wraperrors.Wrap(err, config.ErrorMessageToken401)
	}

	return user, nil
}
func (cli *client) UpdateUser(user *models.User) error {

	return cli.conn.Model(models.User{}).Where(&models.User{ID: user.ID}).Updates(user.ConvertTOMap()).Error
}

func (cli *client) UpdateToken(userID uint, newToken string) error {
	return cli.conn.Model(&models.User{}).Where("id = ?", userID).Update("token", newToken).Error
}

func (cli *client) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return cli.conn.Create(&user).Error
}

func (cli *client) GetUserByToken(token string) (*models.User, error) {
	var user *models.User
	result := cli.conn.Where("token = ?", token).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("No user found with the specified username!")
			return nil, wraperrors.Wrap(result.Error, config.ErrorMessage404)
		} else {
			return nil, wraperrors.Wrap(result.Error, config.ErrorMessage500)
		}
	}
	return user, nil
}
