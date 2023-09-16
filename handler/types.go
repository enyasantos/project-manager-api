package handler

import "time"

type CreateProjectRequest struct {
	Title       string     `json:"title,required"`
	Description string     `json:"description,required"`
	StartDate   time.Time  `json:"start_date,required"`
	EndDate     *time.Time `json:"end_date"`
	Status      string     `json:"status,required"`
	Value       *float64   `json:"value,omitempty"`
}

type PutProjectRequest struct {
	Title       string    `json:"title,required"`
	Description string    `json:"description,required"`
	StartDate   time.Time `json:"start_date,required"`
	EndDate     time.Time `json:"end_date,required"`
	Status      string    `json:"status,required"`
	Value       float64   `json:"value,required"`
}

type PatchProjectRequest struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	Status      *string    `json:"status,omitempty"`
	Value       *float64   `json:"value,omitempty"`
}
