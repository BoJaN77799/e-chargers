package entities

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type RecensionDTO struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	ChargerId uuid.UUID `json:"charger_id"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	Rate      uint      `json:"rate"`
	Toxic     float32   `json:"toxic"`
	Banned    bool      `json:"banned"`
}
