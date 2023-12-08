package route

import (
	"GO-FJ/internal/api/controller"
	"GO-FJ/internal/repository"
	"GO-FJ/internal/usecase"
	"GO-FJ/util"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func NewPostRouter(db *sqlx.DB, timeout time.Duration, group *gin.RouterGroup) {
	pr := repository.NewPostRepository(db)
	pc := controller.PostController{
		PostUsecase: usecase.NewPostUsecase(pr, timeout),
		ImageService: util.Service{
			Destination: "public/",
		},
	}
	post := group.Group("/post")
	{
		post.POST("", pc.CreatePost)
		post.GET("/user", pc.GetPostByUserID)
		post.GET("/titles", pc.GetPostByTitle)
		post.GET("/id", pc.GetPostByID)
		post.PUT("", pc.UpdatePost)
		post.DELETE("", pc.DeletePost)
	}
}
