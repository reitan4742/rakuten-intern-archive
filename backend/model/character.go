package model

type Character struct {
	CharacterID    int    `json:"character_id" gorm:"primary_key"`
	Level          int    `json:"level"`
	CharacterImage string `json:"character"`
}

type CharacterRequest struct {
	Username string `json:"username"`
}

type CharacterResponse struct {
	CharacterImage string `json:"character"`
	Level          int    `json:"level"`
	Exp            int    `json:"exp"`
}
