package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamhi/frontline/api/middlewares"
	"github.com/iamhi/frontline/config"
	"github.com/iamhi/frontline/internal/userhandler"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(context *gin.Context) {
	var request_body loginRequest

	if err := context.ShouldBindJSON(&request_body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	user_details, userhandler_error := userhandler.LoginUser(request_body.Username, request_body.Password)

	if userhandler_error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login", "code": userhandler_error.GetCode()})
		return
	}

	addTokenCookie(context, user_details)

	context.JSON(http.StatusOK, user_details)

	return
}

type createRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func create(context *gin.Context) {
	var request_body createRequest

	if err := context.ShouldBindJSON(&request_body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	user_details, usehandler_error := userhandler.CreateUser(request_body.Username, request_body.Password, request_body.Email)

	if usehandler_error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to create user", "code": usehandler_error.GetCode()})
		return
	}

	addTokenCookie(context, user_details)

	context.JSON(http.StatusOK, user_details)

	return
}

func refreshToken(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token provided"})

		return
	}

	user_details, ok := user_details_obj.(userhandler.UserDetails)

	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token provided"})

		return
	}

	updated_user_details, userhandler_error := userhandler.RefreshToken(user_details)

	if userhandler_error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to refresh token", "code": userhandler_error.GetCode()})
		return
	}

	addTokenCookie(context, updated_user_details)

	context.JSON(http.StatusOK, updated_user_details)

	return
}

func logout(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if exists {
		user_details, ok := user_details_obj.(userhandler.UserDetails)

		if ok {
			userhandler.LogoutUser(user_details)
		}
	}

	removeTokenCookie(context)

	context.JSON(http.StatusOK, gin.H{"message": "ok"})

	return
}

func whoami(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if exists {
		user_details, ok := user_details_obj.(userhandler.UserDetails)

		if ok {
			context.JSON(http.StatusOK, user_details)

			return
		}
	}

	context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token provided"})

	return
}

func addTokenCookie(context *gin.Context, user_details userhandler.UserDetails) {
	context.SetCookie(
		middlewares.AUTHORIZATION_TOKEN_COOKIE,
		user_details.Token,
		config.GetApiCookieConfig().MaxAge,
		config.GetApiCookieConfig().Path,
		config.GetApiCookieConfig().Domain,
		config.GetApiCookieConfig().Secure,
		config.GetApiCookieConfig().HttpOnly)
}

func removeTokenCookie(context *gin.Context) {
	context.SetCookie(
		middlewares.AUTHORIZATION_TOKEN_COOKIE,
		"",
		0,
		config.GetApiCookieConfig().Path,
		config.GetApiCookieConfig().Domain,
		config.GetApiCookieConfig().Secure,
		config.GetApiCookieConfig().HttpOnly)
}

const USER_CONTROLLER_PREFIX = "/user"

func InitializeUserController(parent_router_group *gin.RouterGroup) {
	user_router_group := parent_router_group.Group(USER_CONTROLLER_PREFIX)

	user_router_group.POST("/create", create)
	user_router_group.POST("/login", login)
	user_router_group.POST("/refresh", middlewares.Authorize(), refreshToken)
	user_router_group.POST("/logout", middlewares.AttachUserDetailsMiddleware(), logout)
	user_router_group.GET("/whoami", middlewares.Authorize(), whoami)
}
