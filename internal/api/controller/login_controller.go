package controller

import (
	"GO-FJ/internal/domain"
	"GO-FJ/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginController struct {
	LoginUsecase usecase.LoginUsecase
}

func (lc *LoginController) Login(c *gin.Context) {
	logrus.Info("login request received")

	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		logrus.Errorf("cannot bind a request into login request model: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		logrus.Errorf("user not found with the given email: %s", err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		logrus.Errorf("invalid credentials")
		c.JSON(http.StatusUnauthorized, "invalid credentials")
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, "supersecretkey", 1)
	if err != nil {
		logrus.Errorf("error during creating access token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, "SuperSecretKey", 1)
	if err != nil {
		logrus.Errorf("error during creating refresh token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
