package route

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(db *sql.DB, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignupRouter(db, timeout, publicRouter)
	NewLoginRouter(db, timeout, publicRouter)
}
