package users

import (
	"github.com/aldo-zoop/bookstore-users-api/domain/users"
	"github.com/aldo-zoop/bookstore-users-api/services"
	"github.com/aldo-zoop/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"strconv"

	//"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	//fmt.Println(user)
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	fmt.Println(err.Error())
	//
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//
	//	fmt.Println(err.Error())
	//	return
	//}
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
	//fmt.Println(user)

	//c.String(http.StatusNotImplemented, "Implement me")
}

func SearchUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "Implement me")
}

func GetUser(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("Invalid user")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		restErr := errors.NewNotFoundError("User not found")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "Implement me")
}
