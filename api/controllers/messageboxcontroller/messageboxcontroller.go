package messageboxcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamhi/frontline/api/middlewares"
	"github.com/iamhi/frontline/internal/messageboxhandler"
	"github.com/iamhi/frontline/internal/userhandler"
)

type MessageCreateRequest struct {
	Type    string `json:"type"`
	Subtype string `json:"subtybe"`
	Content string `json:"content"`
}

type createRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func postMyMessage(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	user_details, ok := user_details_obj.(userhandler.UserDetails)

	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	var request_body MessageCreateRequest

	if err := context.ShouldBindJSON(&request_body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	message_dto, messagebox_handler_error := messageboxhandler.PostMyMessage(user_details, request_body.Type, request_body.Subtype, request_body.Content)

	if messagebox_handler_error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": messagebox_handler_error.GetCode()})

		return
	}

	context.JSON(http.StatusOK, message_dto)

	return
}

func getMyMessages(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	user_details, ok := user_details_obj.(userhandler.UserDetails)

	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	messages, messagebox_handler_error := messageboxhandler.GetMyMessages(user_details)

	if messagebox_handler_error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": messagebox_handler_error.GetCode()})

		return
	}

	context.JSON(http.StatusOK, messages)

	return
}

func deleteMessage(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	user_details, ok := user_details_obj.(userhandler.UserDetails)

	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	uuid := context.Param("uuid")

	message, messagebox_handler_error := messageboxhandler.DeleteMessage(user_details, uuid)

	if messagebox_handler_error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": messagebox_handler_error.GetCode()})

		return
	}

	context.JSON(http.StatusOK, message)

	return
}

func postMessage(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	user_details, ok := user_details_obj.(userhandler.UserDetails)

	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	var request_body MessageCreateRequest

	if err := context.ShouldBindJSON(&request_body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	uuid := context.Param("uuid")

	message_dto, messagebox_handler_error := messageboxhandler.PostMessage(user_details, uuid, request_body.Type, request_body.Subtype, request_body.Content)

	if messagebox_handler_error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": messagebox_handler_error.GetCode()})

		return
	}

	context.JSON(http.StatusOK, message_dto)

	return
}

func getBoxMessages(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	user_details, ok := user_details_obj.(userhandler.UserDetails)

	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not-authorized"})

		return
	}

	var request_body MessageCreateRequest

	if err := context.ShouldBindJSON(&request_body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	uuid := context.Param("uuid")

	messages, messagebox_handler_error := messageboxhandler.GetMessages(user_details, uuid)

	if messagebox_handler_error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": messagebox_handler_error.GetCode()})

		return
	}

	context.JSON(http.StatusOK, messages)

	return
}

const MESSAGEBOX_CONTROLLER_PREFIX = "/messagebox"

func InitializeMessageboxController(parent_router_group *gin.RouterGroup) {
	user_router_group := parent_router_group.Group(MESSAGEBOX_CONTROLLER_PREFIX)

	user_router_group.POST("/", middlewares.Authorize(), postMyMessage)
	user_router_group.GET("/", middlewares.Authorize(), getMyMessages)
	user_router_group.DELETE("/message/:uuid/", middlewares.Authorize(), deleteMessage)
	user_router_group.POST("/box/:uuid", middlewares.Authorize(), postMessage)
	user_router_group.GET("/box/:uuid", middlewares.Authorize(), getBoxMessages)
}
