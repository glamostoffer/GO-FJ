package server

import (
	"GO-FJ/internal/controller"
	"GO-FJ/internal/repository"
	"GO-FJ/internal/usecase"
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

func NewSignupRouter(db *sql.DB, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
	}
	group.POST("/signup", sc.Signup)
}
