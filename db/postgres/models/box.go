package models

import "gorm.io/gorm"

type Box struct {
	gorm.Model

	Uuid      string `json:"uuid"`
	OwnerUuid string `json:"owner_uuid"`
}
