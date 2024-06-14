package user

import (
	"github.com/dgrijalva/jwt-go"
	"go-dating-app/helpers"
	"go-dating-app/middlewares"
	"time"
)

type UserService interface {
	Login(request UserReq) (helpers.CommonResponse, error)
}

type userService struct {
	repository UserRepository
}

func NewService(repo UserRepository) UserService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) Login(request UserReq) (response helpers.CommonResponse, err error) {
	userOnDB, err := s.repository.CheckCredentialUser(User{Username: request.Username, Password: request.Password})
	if err != nil {
		return helpers.GenerateErrorWithMessage("internal server error"), nil
	}

	if helpers.IsStringEmpty(userOnDB.Username) {
		return helpers.GenerateErrorWithMessage("invalid username password"), nil
	}

	// Set expiration time for the token
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &middlewares.Claims{
		Username: userOnDB.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the specified algorithm and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(middlewares.JwtKey)
	if err != nil {
		return helpers.GenerateErrorWithMessage("internal server error"), nil
	}

	return helpers.GenerateSuccessWithData("Success login", LoginResp{tokenString}), nil
}
