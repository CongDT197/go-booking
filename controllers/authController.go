package controllers

import (
	"Go_Project/service"
	"Go_Project/utils"
	"Go_Project/utils/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	Username string `json:"user_name" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}

// @Summary get auth
// @ID get-auth
// @Produce json
// @Param user_name body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := utils.Gin{C: c}
	var auth Auth

	valid := validation.Validation{}

	if err := c.ShouldBindJSON(&auth); err != nil {
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, nil)
	}
	a := Auth{Username: auth.Username, Password: auth.Password}
	ok, _ := valid.Valid(&a)

	if !ok {
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, nil)
		return
	}

	authService := service.Auth{Username: auth.Username, Password: auth.Password}
	isExist, err := authService.Check()
	if err != nil {

		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := utils.GenerateToken(auth.Username, auth.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
