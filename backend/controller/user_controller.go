package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
	ru usecase.IReduceUsecase
}

func NewUserController(uu usecase.IUserUsecase, ru usecase.IReduceUsecase) IUserController {
	return &userController{uu, ru}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	userRes := model.UserResponse{}
	userRes.Result = false
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, userRes)
	}
	if err := uc.uu.SignUp(user); err != nil {
		return c.JSON(http.StatusInternalServerError, userRes)
	}
	userReduce := model.UserReduce{}
	userReduce.Username = user.Username
	userReduce.Exp = 0
	userReduce.ReduceScore = 0
	if err := uc.ru.CreateReduce(userReduce); err != nil {
		return c.JSON(http.StatusInternalServerError, userRes)
	}
	userRes.Result = true
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	userRes := model.UserResponse{}
	userRes.Result = false
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, userRes)
	}
	_, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, userRes)
	}
	userRes.Result = true

	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = user.Username
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, userRes)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
