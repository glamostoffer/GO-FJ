package controller

import (
	"GO-FJ/internal/domain"
	"GO-FJ/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type SignupController struct {
	SignupUsecase usecase.SignupUsecase
}

func (sc *SignupController) Signup(c *gin.Context) {
	logrus.Info("signup request received")

	var request domain.SignupRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		logrus.Errorf("cannot bind a request into signup request model: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		logrus.Errorf("error during find user by email method: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if user != (domain.User{}) {
		logrus.Info("attempt to create new user with email which already exists")
		c.JSON(http.StatusConflict, nil)
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		logrus.Errorf("cannot encrypte password: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	request.Password = string(encryptedPassword)

	newUser := domain.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = sc.SignupUsecase.Create(c, &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&newUser, "supersecretkey", 1) // TODO: добавить secret и timeout в cfg
	if err != nil {
		logrus.Errorf("error during creating access token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&newUser, "SuperSecretKey", 1)
	if err != nil {
		logrus.Errorf("error during creating refresh token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusCreated, response)
}
