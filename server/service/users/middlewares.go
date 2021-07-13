package users

import (
	"mini-oa-server/common/auth"
	"mini-oa-server/common/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateContextUserModel(c *gin.Context, this_user_id uint) {
	var thisUserModel User
	if this_user_id != 0 {
		db := database.GetDB()
		db.First(&thisUserModel, this_user_id)
	}

	c.Set("this_user_id", this_user_id)
	c.Set("this_user_model", thisUserModel)
}

func ResetContextUserModel(c *gin.Context) {
	UpdateContextUserModel(c, 0)
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ResetContextUserModel(c)

		jwtInst := auth.GetJwtInstance()
		token, err := jwtInst.ParseToken(c)

		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}

		if claims, valid := jwtInst.VerifyTokenPayload(token); valid {
			UpdateContextUserModel(c, claims.ID)
		} else {
			if auto401 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			return
		}
	}
}
