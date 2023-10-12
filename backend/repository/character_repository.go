package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type ICharacterRepository interface {
	GetCharacterByCharacterId(character *model.Character, characterId int) error
	GetCharacterByLevel(character *model.Character, level int) error
	CreateCharacter(character *model.Character) error
}

type characterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) ICharacterRepository {
	return &characterRepository{db}
}

func (cr *characterRepository) GetCharacterByCharacterId(character *model.Character, characterId int) error {
	if err := cr.db.Where("character_id = ?", characterId).First(character).Error; err != nil {
		return err
	}
	return nil
}

func (cr *characterRepository) GetCharacterByLevel(character *model.Character, level int) error {
	if err := cr.db.Where("level = ?", level).First(character).Error; err != nil {
		return err
	}
	return nil
}

func (cr *characterRepository) CreateCharacter(character *model.Character) error {
	if err := cr.db.Create(character).Error; err != nil {
		return err
	}
	return nil
}
