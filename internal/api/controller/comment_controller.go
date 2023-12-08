package controller

import (
	"GO-FJ/internal/domain"
	"GO-FJ/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CommentController struct {
	CommentUsesase usecase.CommentUsecase
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	logrus.Info("create comment request received")
	var comment domain.Comment
	err := c.ShouldBind(&comment)
	if err != nil {
		logrus.Errorf("cannot bind request body: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	comment.UserID = int64(c.GetFloat64("user-id"))
	err = cc.CommentUsesase.Create(c, &comment)
	if err != nil {
		logrus.Errorf("cannot create comment: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (cc *CommentController) GetCommentByID(c *gin.Context) {
	logrus.Info("get comment by id request received")
	queryParams := c.Request.URL.Query()
	id := queryParams.Get("id")

	comment, err := cc.CommentUsesase.GetByID(c, id)
	if err != nil {
		logrus.Errorf("cannot get comment: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (cc *CommentController) GetCommentByPostID(c *gin.Context) {
	logrus.Info("get comment by id request received")
	queryParams := c.Request.URL.Query()
	id := queryParams.Get("id")

	comments, err := cc.CommentUsesase.GetByPostID(c, id)
	if err != nil {
		logrus.Errorf("cannot get comments: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	logrus.Info("update comment by id request received")
	queryParams := c.Request.URL.Query()
	id := queryParams.Get("id")

	userID := c.GetFloat64("user-id")
	comment, err := cc.CommentUsesase.GetByID(c, id)
	if err != nil {
		logrus.Errorf("cannot get comment: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if comment.UserID != int64(userID) {
		logrus.Errorf("trying to update someone else's comment")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	err = c.ShouldBind(&comment)
	if err != nil {
		logrus.Errorf("cannot bind comment: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = cc.CommentUsesase.UpdateComment(c, comment)
	if err != nil {
		logrus.Errorf("cannot update comment: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	logrus.Info("delete comment by id request received")
	queryParams := c.Request.URL.Query()
	id := queryParams.Get("id")

	userID := c.GetFloat64("user-id")
	comment, err := cc.CommentUsesase.GetByID(c, id)
	if err != nil {
		logrus.Errorf("cannot get comment: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if comment.UserID != int64(userID) {
		logrus.Errorf("trying to delete someone else's comment")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	err = cc.CommentUsesase.DeleteComment(c, id)
	if err != nil {
		logrus.Errorf("cannot update comment: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
