package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type WorkReportData struct {
	WorkReport WorkReport `json:"work_report"`
	Message    string     `json:"message"`
}

// CreateWorkReport example
type CreateWorkReport struct {
	Worker       string  `json:"worker" example:"John Doe"`
	Work         string  `json:"work" example:"Hammering"`
	WorkQuantity float64 `json:"work_quantity,omitempty" example:"10.0"`
	QuantityUnit string  `json:"quantity_unit,omitempty" example:"Kg"`
	SiteID       uint64  `json:"site_id,omitempty" example:"1"`
}

// WorkReport example
type WorkReport struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at,omitempty" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" example:"2020-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	CreateWorkReport
}

// example
var (
	ErrWorkerInvalid = errors.New("Worker is invalid or empty")
	ErrWorkInvalid   = errors.New("Work is invalid or empty")
	ErrQuantityUnit  = errors.New("Quantity unit is invalid or empty")
	ErrWorkQuantity  = errors.New("Work quantity is invalid or empty")
)

// Validation validates CreateWorkReport
func (w *CreateWorkReport) Validation() error {
	switch {
	case len(w.Worker) == 0 || w.Worker == "":
		return ErrWorkerInvalid
	case len(w.Work) == 0 || w.Work == "":
		return ErrWorkInvalid
	case len(w.QuantityUnit) == 0 || w.QuantityUnit == "":
		return ErrQuantityUnit
	case len(w.Work) == 0 || w.WorkQuantity == 0:
		return ErrWorkQuantity
	default:
		return nil
	}
}
