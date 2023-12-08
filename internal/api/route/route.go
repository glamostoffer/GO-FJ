package route

import (
	"GO-FJ/internal/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func Setup(db *sqlx.DB, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignupRouter(db, timeout, publicRouter)
	NewLoginRouter(db, timeout, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuth("supersecretkey"))
	NewPostRouter(db, timeout, protectedRouter)
	NewCommentRouter(db, timeout, protectedRouter)
}
