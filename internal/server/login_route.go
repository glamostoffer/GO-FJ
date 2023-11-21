package server

import (
	"GO-FJ/internal/controller"
	"GO-FJ/internal/repository"
	"GO-FJ/internal/usecase"
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

func NewLoginRouter(db *sql.DB, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
	}
	group.POST("/login", lc.Login)
}
