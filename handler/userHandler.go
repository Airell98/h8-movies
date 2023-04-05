package handler

import (
	"h8-movies/dto"
	"h8-movies/pkg/errs"
	"h8-movies/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

// UserRegister godoc
// @Tags users
// @Description Create New User Data
// @ID create-new-user
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewUserRequest true "request body json"
// @Success 201 {object} dto.NewUserResponse
// @Router /users/register [post]
func (uh *userHandler) Register(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.CreateNewUser(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)

}

// UserLogin godoc
// @Tags users
// @Description User Sign In
// @ID user-sign-in
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewUserRequest true "request body json"
// @Success 200 {object} dto.LoginResponse
// @Router /users/login [post]
func (uh *userHandler) Login(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.Login(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
