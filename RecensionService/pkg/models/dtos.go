package models

type RecensionDTO struct {
	Id        uint    `json:"id"`
	Username  string  `json:"username"`
	ChargerId uint    `json:"charger_id"`
	Date      uint64  `json:"date"`
	Content   string  `json:"content"`
	Rate      uint    `json:"rate"`
	Toxic     float32 `json:"toxic"`
	Banned    bool    `json:"banned"`
}
