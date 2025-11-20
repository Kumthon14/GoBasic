package Controllers

import (
	"GoPractice/Models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// @Summary Get All User Data
// @Description Get All User Data
// @ID GetAllUserData
// @Tags User Data
// @Success 200 {object} []Models.User "Success"
// @Failure 400 {string} string "Error"
// @response 401 {string} string "Unauthorized"
// @Router /user-api/user [GET]
// @security ApiKeyAuth
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// @Summary Register User
// @Description Register New User
// @ID RegisterUser
// @Tags User
// @Param EnterDetails body Models.UserRegister true "Register"
// @Accept json
// @Success 200 {object} Models.User "Success"
// @Failure 400 {string} string "Error"
// @Router /auth-api/register [POST]
func CreateUser(c *gin.Context) {
	var user Models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Models.CreateUser(&user)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUserById(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")

	err := Models.GetUserById(&user, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")

	err := Models.GetUserById(&user, id)

	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}

	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}

}

// @Summary Login User
// @Description Login User
// @ID LoginUser
// @Tags User
// @Param EnterDetails body Models.UserLogin true "Login"
// @Accept json
// @Success 200 {object} Models.UserLoginRes "Success"
// @Failure 400 {string} string "Error"
// @Router /auth-api/login [POST]
func Login(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.Login(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		})

		ss, err := token.SignedString([]byte("MySignature"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		userRes := Models.UserLoginRes{
			Name:     user.Name,
			Email:    user.Email,
			Phone:    user.Phone,
			Address:  user.Address,
			Username: user.Username,
			Token:    ss,
		}

		c.JSON(http.StatusOK, userRes)
	}

}
