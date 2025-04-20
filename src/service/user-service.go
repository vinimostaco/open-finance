package service

import (
	"errors"

	"github.com/vinimostaco/open-finance/src/config"
	"github.com/vinimostaco/open-finance/src/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(input model.RegisterUserInput) (*model.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &model.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: string(hashedPassword),
    }

    if err := config.DB.Create(user).Error; err != nil {
        return nil, err
    }

    return user, nil
}

func AuthenticateUser(email, password string) (*model.User, error) {
    var user model.User
    if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, errors.New("usuário não encontrado")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("senha incorreta")
    }

    return &user, nil
}