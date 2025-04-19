package websocket

import (
	"real-chat/service"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	messageService *service.MessageService
}

func (h *MessageHandler) NewMesssageHandler (*gin.Context) {

}

