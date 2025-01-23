package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	"log"
	"otzovik-back/internal/domain"
)

type WebsocketHandler struct {
	service  domain.ChatService
	upgrader websocket.Upgrader
}

func NewWebsocketHandler(service domain.ChatService, upgrader websocket.Upgrader) *WebsocketHandler {
	return &WebsocketHandler{service, upgrader}
}

func (w *WebsocketHandler) InitWebsocket(c *gin.Context) {
	conn, err := w.upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(conn)

	userId := c.Query("id")
	log.Println("User ID:", userId)

	for {
		mt, message, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("Received message: %s\n", message)

		responseMsg := fmt.Sprintf("Echo %s", message)

		err = conn.WriteMessage(mt, []byte(responseMsg))

		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}
