package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model

	Uuid      string `json:"uuid"`
	OwnerUuid string `json:"owner_uuid"`
	BoxUuid   string `json:"box_uuid"`
	Type      string `json:"type"`
	Subtype   string `json:"subtype"`
	Content   string `json:"content"`
}
