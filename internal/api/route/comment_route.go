package route

import (
	"GO-FJ/internal/api/controller"
	"GO-FJ/internal/repository"
	"GO-FJ/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func NewCommentRouter(db *sqlx.DB, timeout time.Duration, group *gin.RouterGroup) {
	cr := repository.NewCommentRepository(db)
	cc := controller.CommentController{
		CommentUsesase: usecase.NewCommentUsecase(cr, timeout),
	}
	comment := group.Group("/comment")
	{
		comment.POST("", cc.CreateComment)
		comment.GET("", cc.GetCommentByID)
		comment.GET("/post", cc.GetCommentByPostID)
		comment.PUT("", cc.UpdateComment)
		comment.DELETE("", cc.DeleteComment)
	}
}
