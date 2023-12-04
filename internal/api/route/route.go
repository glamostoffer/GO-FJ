package route

import (
	"GO-FJ/internal/api/middleware"
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(db *sql.DB, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignupRouter(db, timeout, publicRouter)
	NewLoginRouter(db, timeout, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuth("supersecretkey"))
	NewPostRouter(db, timeout, protectedRouter)
}
