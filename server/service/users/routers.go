package users

import (
	"log"
	"mini-oa-server/common/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouters(r *gin.RouterGroup) {
	r.POST("/register", userRegisterRoute)
	r.POST("/login", userLoginRoute)
}

func userLoginRoute(c *gin.Context) {
	var lv LoginValidator
	if err := lv.Bind(c); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error in posted body",
		})
		return
	}

	userModel, err := FindOneUser(&User{Username: lv.User.Username})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "cannot find this user",
		})
		return
	}

	if userModel.CheckPassword(lv.User.Password) != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "The user name or the password cannot match the record",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"user_id":  userModel.ID,
			"username": userModel.Username,
			"token":    auth.GenJwtToken(userModel.ID),
		},
	})
}

func userRegisterRoute(c *gin.Context) {
	var uv UserValidator
	if err := uv.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error in posted body",
		})
		return
	}

	if err := uv.UserModel.Save(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register succeeded",
	})
}
