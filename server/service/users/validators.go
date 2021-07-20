package users

import "github.com/gin-gonic/gin"

type UserValidator struct {
	User struct {
		Username string `json:"username" binding:"required,alphanum,min=4,max=255"`
		Email    string `son:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,alphanum,min=8,max=255"`
	} `json:"user"`

	UserModel User `json:"-"`
}

func (uv *UserValidator) Bind(c *gin.Context) error {
	err := c.ShouldBind(&uv)
	if err != nil {
		return err
	}

	uv.UserModel.Username = uv.User.Username
	uv.UserModel.Email = uv.User.Email
	uv.UserModel.SetPassword(uv.User.Password)

	return nil
}

type LoginValidator struct {
	User struct {
		Username string `json:"username" binding:"required,alphanum,min=4,max=255"`
		Password string `json:"password" binding:"required,alphanum,min=8,max=255"`
	} `json:"user"`

	IsCheck bool `json:"is_check"`
}

func (lv *LoginValidator) Bind(c *gin.Context) error {
	err := c.ShouldBind(&lv)
	if err != nil {
		return err
	}
	return nil
}
