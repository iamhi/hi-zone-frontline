package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iamhi/frontline/api/controllers/messageboxcontroller"
	"github.com/iamhi/frontline/api/controllers/usercontroller"
)

const SERVICE_PREFIX = "/hi-zone-api/frontline"

func initialize(gin_engine *gin.Engine) {
	root_group := gin_engine.Group(SERVICE_PREFIX)

	usercontroller.InitializeUserController(root_group)
	messageboxcontroller.InitializeMessageboxController(root_group)
}
