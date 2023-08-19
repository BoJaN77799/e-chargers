package entities

import (
	uuid "github.com/satori/go.uuid"
	"recension_service/pkg/entities/base"
	"time"
)

type Recension struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key;type:uuid"`
	UserId    uuid.UUID `json:"user_id" gorm:"not-null;type:uuid"`
	ChargerId uuid.UUID `json:"charger_id" gorm:"not-null;type:uuid"`
	Date      time.Time `json:"date"  gorm:"not-null;default:CURRENT_TIMESTAMP"`
	Content   string    `json:"content" gorm:"not-null"`
	Rate      uint      `json:"rate" gorm:"not-null"`
	Toxic     float32   `json:"toxic" gorm:"not-null"`
	Banned    bool      `json:"banned" gorm:"not-null"`
	base.Base
}

func (recension *Recension) ToDTO() RecensionDTO {
	return RecensionDTO{
		Id:        recension.Id,
		UserId:    recension.UserId,
		ChargerId: recension.ChargerId,
		Date:      recension.Date,
		Content:   recension.Content,
		Rate:      recension.Rate,
		Toxic:     recension.Toxic,
		Banned:    recension.Banned,
	}
}
