package router

import (
	"net/http"

	"github.com/Fakorede/gin-app/controllers"
	"github.com/Fakorede/gin-app/dto"
	"github.com/gin-gonic/gin"
)

type AuthApis struct {
	loginController controllers.LoginController
}

func NewAuthAPIs(loginController controllers.LoginController) *AuthApis {
	return &AuthApis{
		loginController: loginController,
	}
}

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} dto.JWT
// @Failure 401 {object} dto.Response
// @Router /auth/token [post]
func (api *AuthApis) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &dto.JWT{
			Token: token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{
			Message: "Invalid credentials",
		})
	}
}