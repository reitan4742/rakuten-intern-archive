package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	receiptRepository := repository.NewReceiptRepository(db)
	reduceRepository := repository.NewUserReduceRepository(db)
	allReduceRepository := repository.NewAllReduceRepository(db)
	characterRepository := repository.NewCharacterRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	receiptUsecase := usecase.NewReceiptUsecase(receiptRepository)
	reduceUsecase := usecase.NewReduceUsecase(reduceRepository, allReduceRepository)
	characterUsecase := usecase.NewCharacterUsecase(characterRepository)

	userController := controller.NewUserController(userUsecase, reduceUsecase)
	receiptController := controller.NewReceiptController(receiptUsecase, reduceUsecase)
	reduceController := controller.NewReduceController(reduceUsecase)
	characterController := controller.NewCharacterController(characterUsecase, reduceUsecase)

	e := router.NewRouter(userController, receiptController, reduceController, characterController)
	e.Logger.Fatal(e.Start(":8080"))
}
