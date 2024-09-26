package controller

import (
	"context"
	"net/http"

	"belazap.com/assistant/assistant"
	"github.com/gin-gonic/gin"
)

type chatController struct {
	ai assistant.AssistantClient
}

type ChatController interface {
	postAnswer()
}

type PostChatRequestDTO struct {
	Prompt string `json:"prompt"`
}

func NewChatController(r *gin.RouterGroup) {
	controller := &chatController{
		ai: assistant.New(context.Background()),
	}
	r.POST("/", controller.postChat)
}

func (cc *chatController) postChat(c *gin.Context) {
	body := PostChatRequestDTO{}

	if err := c.ShouldBind(&body); err == nil {
		answer, err := cc.ai.Answer(body.Prompt)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"answer": answer,
			})
		}
	}
}
