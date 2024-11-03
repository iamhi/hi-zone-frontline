package messageboxhandler

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/iamhi/frontline/db/postgres"
	"github.com/iamhi/frontline/db/postgres/models"
	"github.com/iamhi/frontline/internal/errors"
	"github.com/iamhi/frontline/internal/userhandler"
)

type MessageDto struct {
	Uuid      string    `json:"uuid"`
	OwnerUuid string    `json:"owner_uuid"`
	BoxUuid   string    `json:"box_uuid"`
	Type      string    `json:"title"`
	Subtype   string    `json:"subtype"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type BoxDto struct {
	Uuid      string `json:"uuid"`
	OwnerUuid string `json:"owner_uuid"`
}

func GetMessages(user_details userhandler.UserDetails, box_uuid string) ([]MessageDto, errors.MessageboxHandlerError) {
	var messageEntities []models.Message
	messageDtos := []MessageDto{}

	referenced_box := findBoxByUuid(box_uuid)

	if referenced_box.ID == 0 {
		var messagebox_handler_error errors.MessageboxHandlerError

		referenced_box, messagebox_handler_error = createBox(user_details.Uuid)

		if messagebox_handler_error != nil {
			return []MessageDto{}, messagebox_handler_error
		}
	}

	if referenced_box.OwnerUuid != user_details.Uuid &&
		!strings.Contains(user_details.Roles, userhandler.USER_ROLE_SERVICE) {
		return []MessageDto{}, &errors.MessageboxAccessNotAllowedError{}
	}

	postgres.Db.Model(&models.Message{}).Where("box_uuid=?", referenced_box.Uuid).Find(&messageEntities)

	for _, message := range messageEntities {
		messageDtos = append(messageDtos, messageEntityToDto(message))
	}

	return messageDtos, nil
}

func GetMyMessages(user_details userhandler.UserDetails) ([]MessageDto, errors.MessageboxHandlerError) {
	referenced_box := findBoxByOwnerUuid(user_details.Uuid)

	return GetMessages(user_details, referenced_box.Uuid)
}

func PostMessage(
	user_details userhandler.UserDetails,
	box_uuid string,
	message_type string,
	message_subtype string,
	message_content string) (MessageDto, errors.MessageboxHandlerError) {
	var saved_message_entity models.Message
	message_uuid := uuid.New().String()

	referenced_box := findBoxByUuid(box_uuid)

	if referenced_box.ID == 0 {
		return MessageDto{}, &errors.MessageboxNotFoundError{}
	}

	if referenced_box.OwnerUuid != user_details.Uuid &&
		!strings.Contains(user_details.Roles, userhandler.USER_ROLE_SERVICE) {
		return MessageDto{}, &errors.MessageboxAccessNotAllowedError{}
	}

	postgres.Db.Create(&models.Message{
		Uuid:      message_uuid,
		OwnerUuid: user_details.Uuid,
		BoxUuid:   referenced_box.Uuid,
		Type:      message_type,
		Subtype:   message_subtype,
		Content:   message_content,
	})

	postgres.Db.Model(&models.Message{}).Where("uuid=?", message_uuid).Find(&saved_message_entity)

	return messageEntityToDto(saved_message_entity), nil
}

func PostMyMessage(
	user_details userhandler.UserDetails,
	message_type string,
	message_subtype string,
	message_content string) (MessageDto, errors.MessageboxHandlerError) {
	referenced_box := findBoxByOwnerUuid(user_details.Uuid)

	return PostMessage(user_details, referenced_box.Uuid, message_type, message_subtype, message_content)
}

func DeleteMessage(
	user_details userhandler.UserDetails,
	message_uuid string,
) (MessageDto, errors.MessageboxHandlerError) {
	referenced_box := findBoxByUuid(user_details.Uuid)
	message_entity := findMessageByUuid(message_uuid)

	if referenced_box.Uuid != message_entity.BoxUuid &&
		!strings.Contains(user_details.Roles, userhandler.USER_ROLE_SERVICE) {
		return MessageDto{}, &errors.MessageboxAccessNotAllowedError{}
	}

	if message_entity.ID != 0 {
		postgres.Db.Delete(&models.Message{}, "uuid=?", message_uuid)

		return MessageDto{
			Uuid: message_entity.Uuid,
		}, nil
	}

	return MessageDto{}, nil
}

func createBox(owner_uuid string) (models.Box, errors.MessageboxHandlerError) {
	var existing_box models.Box
	var created_box models.Box

	postgres.Db.Model(&models.Box{}).Where("owner_uuid=?", owner_uuid).Find(&existing_box)

	if existing_box.ID != 0 {
		return models.Box{}, &errors.MessageboxAlreadyExistsError{}

	}

	box_uuid := uuid.New().String()

	postgres.Db.Create(&models.Box{
		Uuid:      box_uuid,
		OwnerUuid: owner_uuid,
	})

	postgres.Db.Model(&models.Box{}).Where("owner_uuid=?", owner_uuid).Find(&created_box)

	return created_box, nil
}

func findBoxByOwnerUuid(owner_uuid string) models.Box {
	var referenced_box models.Box

	postgres.Db.Model(&models.Box{}).Where("owner_uuid=?", owner_uuid).Find(&referenced_box)

	return referenced_box
}

func findBoxByUuid(box_uuid string) models.Box {
	var referenced_box models.Box

	postgres.Db.Model(&models.Box{}).Where("uuid=?", box_uuid).Find(&referenced_box)

	return referenced_box
}

func findMessageByUuid(uuid string) models.Message {
	var message_entity models.Message

	postgres.Db.Model(&models.Message{}).Where("uuid=?", uuid).Find(&message_entity)

	return message_entity
}

func messageEntityToDto(entity models.Message) MessageDto {
	return MessageDto{
		Uuid:      entity.Uuid,
		OwnerUuid: entity.OwnerUuid,
		BoxUuid:   entity.BoxUuid,
		Type:      entity.Type,
		Subtype:   entity.Subtype,
		Content:   entity.Content,
		CreatedAt: entity.CreatedAt,
	}
}
