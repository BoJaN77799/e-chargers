package models

type RecensionDTO struct {
	Username  string `json:"username"`
	ChargerId uint   `json:"charger_id"`
	Date      uint64 `json:"date"`
	Content   string `json:"content"`
	Rate      uint   `json:"rate"`
}
