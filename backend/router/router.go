package router

import (
	"backend/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, rcc controller.IReceiptController,
	rdc controller.IReduceController, cc controller.ICharacterController) *echo.Echo {

	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)
	e.POST("/home", cc.Home)
	e.POST("/receipt", rcc.GetReceipt)
	e.POST("/receiptresult", rcc.GetReceiptResult)
	e.POST("/history", rcc.GetHistory)
	e.POST("/foodlossreduce/person", rdc.GetUserReduce)
	e.POST("/foodlossreduce/all", rdc.GetAllReduce)
	return e
}
