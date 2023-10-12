package repository

import (
	"backend/model"
	"errors"

	"gorm.io/gorm"
)

type IUserReduceRepository interface {
	GetUserReduceByUsername(userReduce *model.UserReduce, username string) error
	CreateUserReduce(userReduce *model.UserReduce) error
	UpdateUserReduce(userReduce *model.UserReduce) error
}

type IAllReduceRepository interface {
	GetAllReduce(allReduce *model.AllReduce) error
	UpdateAllReduce(allReduce *model.AllReduce) error
}

type userReduceRepository struct {
	db *gorm.DB
}

type allReduceRepository struct {
	db *gorm.DB
}

func NewUserReduceRepository(db *gorm.DB) IUserReduceRepository {
	return &userReduceRepository{db}
}

func NewAllReduceRepository(db *gorm.DB) IAllReduceRepository {
	return &allReduceRepository{db}
}

func (urr *userReduceRepository) GetUserReduceByUsername(userReduce *model.UserReduce, username string) error {
	if err := urr.db.Where("username = ?", username).First(userReduce).Error; err != nil {
		return err
	}
	return nil
}

func (urr *userReduceRepository) CreateUserReduce(userReduce *model.UserReduce) error {
	if err := urr.db.Create(userReduce).Error; err != nil {
		return err
	}
	return nil
}

func (urr *userReduceRepository) UpdateUserReduce(userReduce *model.UserReduce) error {
	result := urr.db.Model(&userReduce).Where("username = ?", userReduce.Username).
		Updates(map[string]interface{}{"exp": userReduce.Exp, "reduce_score": userReduce.ReduceScore})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (arr *allReduceRepository) GetAllReduce(allReduce *model.AllReduce) error {
	if err := arr.db.Where("all_reduce_id = ?", 1).First(allReduce).Error; err != nil {
		return err
	}
	return nil
}

func (arr *allReduceRepository) UpdateAllReduce(allReduce *model.AllReduce) error {
	// get 1 row 1 column
	err := arr.db.Model(&allReduce).Where("all_reduce_id = ?", 1).Update("all_reduce", allReduce.AllReduce)
	if err != nil {
		return err.Error
	}
	return nil
}
