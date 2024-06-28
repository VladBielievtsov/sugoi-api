package types

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	ID          *uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"id,omitempty"`
	ImageURL    string     `gorm:"type:varchar(255);unique;not null" json:"image_url,omitempty"`
	ImageSize   int        `gorm:"not null" json:"image_size,omitempty"`
	ImageWidth  int        `gorm:"not null" json:"image_width,omitempty"`
	ImageHeight int        `gorm:"not null" json:"image_height,omitempty"`
	Source      string     `gorm:"type:varchar(255)" json:"source,omitempty"`
	Tags        []Tag      `gorm:"many2many:image_tags" json:"tags,omitempty"`
	CreatedAt   time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Tag struct {
	ID          *uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"id,omitempty"`
	Name        string     `gorm:"type:varchar(255);unique;not null" json:"name,omitempty"`
	Description string     `gorm:"type:varchar(255);not null" json:"description,omitempty"`
	Images      []Image    `gorm:"many2many:image_tags" json:"images,omitempty"`
}

type CreateTagBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// Characters  []Character `gorm:"many2many:image_characters" json:"characters,omitempty"`

// type Character struct {
// 	ID          *uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"id,omitempty"`
// 	Name        string     `gorm:"type:varchar(255);unique;not null" json:"name,omitempty"`
// 	Description string     `gorm:"type:varchar(255);not null" json:"description,omitempty"`
// 	Gender      string     `gorm:"type:varchar(255);not null" json:"gender,omitempty"`
// 	Species     string     `gorm:"type:varchar(255);not null" json:"species,omitempty"`
// 	Images      []Image    `gorm:"many2many:image_tags" json:"images,omitempty"`
// }
