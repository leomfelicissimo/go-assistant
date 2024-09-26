package rest

import (
	"belazap.com/assistant/controller"
	"github.com/gin-gonic/gin"
)

func StartHttp() {
	r := gin.Default()
	a := r.Group("/chats")
	{
		controller.NewChatController(a)
	}
	r.Run()
}
