package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tomyda/bookstore_users-api/domain/users"
	"github.com/tomyda/bookstore_users-api/services"
	"github.com/tomyda/bookstore_users-api/utils/errors"
)

//CreateUser func
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetUser func
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_ID"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getError := services.GetUser(userID)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, user)
}
