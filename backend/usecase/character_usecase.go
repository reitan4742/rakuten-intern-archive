package usecase

import (
	"backend/model"
	"backend/repository"
)

type ICharacterUsecase interface {
	GetCharacterByLevel(level int) (string, error)
}

type characterUsecase struct {
	cr repository.ICharacterRepository
}

func NewCharacterUsecase(cr repository.ICharacterRepository) ICharacterUsecase {
	return &characterUsecase{cr}
}

func (cu *characterUsecase) GetCharacterByLevel(level int) (string, error) {
	character := model.Character{}
	if err := cu.cr.GetCharacterByLevel(&character, level); err != nil {
		return "", err
	}
	return character.CharacterImage, nil
}
