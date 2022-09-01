package models

import (
	"github.com/jinzhu/gorm"
)

type Recension struct {
	gorm.Model
	Username  string  `json:"username" gorm:"not-null"`
	ChargerId uint    `json:"charger_id" gorm:"not-null"`
	Date      uint64  `json:"date"  gorm:"not-null"`
	Content   string  `json:"content" gorm:"not-null"`
	Rate      uint    `json:"rate" gorm:"not-null"`
	Toxic     float32 `json:"toxic" gorm:"not-null"`
}

func (recension *Recension) ToDTO() RecensionDTO {
	return RecensionDTO{
		Username:  recension.Username,
		ChargerId: recension.ChargerId,
		Date:      recension.Date,
		Content:   recension.Content,
		Rate:      recension.Rate,
		Toxic:     recension.Toxic,
	}
}

type CancelRecension struct {
	Username  string `json:"username"`
	ChargerId uint   `json:"charger_id"`
}
