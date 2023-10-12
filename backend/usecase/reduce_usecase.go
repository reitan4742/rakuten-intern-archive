package usecase

import (
	"backend/model"
	"backend/repository"
)

type IReduceUsecase interface {
	CreateReduce(reduce model.UserReduce) error
	AddUserScore(username string, exp int, reduceScore int) error
	AddAllScore(allReduce int) error
	GetUserReduceByUsername(username string) (int, error)
	GetUserExpByUsername(username string) (int, error)
	GetAllReduce() (int, error)
}

type reduceUsecase struct {
	urr repository.IUserReduceRepository
	arr repository.IAllReduceRepository
}

func NewReduceUsecase(urr repository.IUserReduceRepository, arr repository.IAllReduceRepository) IReduceUsecase {
	return &reduceUsecase{urr, arr}
}

func (ru *reduceUsecase) CreateReduce(reduce model.UserReduce) error {
	if err := ru.urr.CreateUserReduce(&reduce); err != nil {
		return err
	}
	return nil
}

func (ru *reduceUsecase) AddUserScore(username string, exp int, reduceScore int) error {
	userReduce := model.UserReduce{}
	if err := ru.urr.GetUserReduceByUsername(&userReduce, username); err != nil {
		return err
	}
	userReduce.Exp += exp
	userReduce.ReduceScore += reduceScore
	if err := ru.urr.UpdateUserReduce(&userReduce); err != nil {
		return err
	}
	return nil
}

func (ru *reduceUsecase) AddAllScore(allReduce int) error {
	allReduceModel := model.AllReduce{}
	if err := ru.arr.GetAllReduce(&allReduceModel); err != nil {
		return err
	}
	allReduceModel.AllReduce += allReduce
	if err := ru.arr.UpdateAllReduce(&allReduceModel); err != nil {
		return err
	}
	return nil
}

func (ru *reduceUsecase) GetUserReduceByUsername(username string) (int, error) {
	userReduce := model.UserReduce{}
	if err := ru.urr.GetUserReduceByUsername(&userReduce, username); err != nil {
		return 0, err
	}
	return userReduce.ReduceScore, nil
}

func (ru *reduceUsecase) GetUserExpByUsername(username string) (int, error) {
	userReduce := model.UserReduce{}
	if err := ru.urr.GetUserReduceByUsername(&userReduce, username); err != nil {
		return 0, err
	}
	return userReduce.Exp, nil
}

func (ru *reduceUsecase) GetAllReduce() (int, error) {
	allReduce := model.AllReduce{}
	if err := ru.arr.GetAllReduce(&allReduce); err != nil {
		return 0, err
	}
	return allReduce.AllReduce, nil
}
