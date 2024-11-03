package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamhi/frontline/internal/userhandler"
)

const USER_DETAILS = "user_details"
const AUTHORIZATION_TOKEN_COOKIE = "Authorization-Token"

func AttachUserDetailsMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization_cookie, err := context.Cookie(AUTHORIZATION_TOKEN_COOKIE)
		authorization := context.GetHeader("Authorization")

		if err != nil && authorization == "" {
			context.Next()
			return
		} else if authorization != "" {
			authorization_cookie = authorization
		}

		user_details, userhandler_error := userhandler.GetUserDetails(authorization_cookie)

		if userhandler_error == nil {
			context.Set(USER_DETAILS, user_details)
		}

		context.Next()
	}
}

func Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization_cookie, err := context.Cookie(AUTHORIZATION_TOKEN_COOKIE)
		authorization := context.GetHeader("Authorization")

		if err != nil && authorization == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			context.Abort()

			return
		}

		if authorization != "" {
			authorization_cookie = authorization
		}

		user_details, userhandler_error := userhandler.GetUserDetails(authorization_cookie)

		if userhandler_error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": userhandler_error.GetCode()})
			context.Abort()

			return
		}

		context.Set(USER_DETAILS, user_details)

		context.Next()
	}
}
