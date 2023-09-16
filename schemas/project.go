package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
	Value       float64   `json:"value,omitempty"`
}

type ProjectResponse struct {
	ID          uuid.UUID  `gorm:"type:uuid;"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deteledAt,omitempty"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Status      *string    `json:"status"`
	Value       *float64   `json:"value,omitempty"`
}

func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	project.ID = uuid.New()
	return
}
