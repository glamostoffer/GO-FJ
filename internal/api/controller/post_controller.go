package controller

import (
	"GO-FJ/internal/domain"
	"GO-FJ/internal/usecase"
	"GO-FJ/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

type PostController struct {
	PostUsecase  usecase.PostUsecase
	ImageService util.Service
}

func (pc *PostController) CreatePost(c *gin.Context) {
	logrus.Info("create post request received")
	form, err := c.MultipartForm()
	if err != nil {
		logrus.Errorf("cannot get images from form: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	files := form.File["images"]
	paths, err := pc.ImageService.SaveImages(files)
	if err != nil {
		logrus.Errorf("cannot save image: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userID := c.GetFloat64("user-id")
	text := form.Value["text"]
	title := form.Value["title"]

	log.Printf("userID: %s\n", userID)

	post := domain.Post{
		UserID:    int64(userID),
		Title:     title[0],
		Text:      text[0],
		Images:    paths,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = pc.PostUsecase.Create(c, &post)
	if err != nil {
		logrus.Errorf("error during creating post: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (pc *PostController) GetPostByTitle(c *gin.Context) {
	logrus.Info("get posts by title request received")
	queryParams := c.Request.URL.Query()
	title := queryParams.Get("title")

	posts, err := pc.PostUsecase.GetByTitle(c, title)
	if err != nil {
		logrus.Errorf("cannot get posts with title %s: %s", title, err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (pc *PostController) GetPostByUserID(c *gin.Context) {
	logrus.Info("get posts by user id request received")
	queryParams := c.Request.URL.Query()
	id := queryParams.Get("id")

	posts, err := pc.PostUsecase.GetByUserID(c, id)
	if err != nil {
		logrus.Errorf("cannot get posts where author's id is  %s: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (pc *PostController) GetPostByID(c *gin.Context) {
	logrus.Info("get post by id request received")
	queryParams := c.Request.URL.Query()
	id := queryParams.Get("id")

	post, err := pc.PostUsecase.GetByID(c, id)
	if err != nil {
		logrus.Errorf("cannot get posts where id is %s: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	logrus.Info("update post request received")
	queryParams := c.Request.URL.Query()
	id := queryParams.Get("id")
	userID := c.GetFloat64("user-id")

	post, err := pc.PostUsecase.GetByID(c, id)
	if err != nil {
		logrus.Errorf("cannot get posts where id is %s: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if post.UserID != int64(userID) {
		logrus.Errorf("trying to update someone else's post")
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	err = c.ShouldBind(&post)
	if err != nil {
		logrus.Errorf("cannot bind a request into post update model: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = pc.PostUsecase.UpdatePost(c, post)
	if err != nil {
		logrus.Errorf("cannot update post: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
