package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	"log"
)

type WebsocketHandler struct {
	upgrader websocket.Upgrader
}

func NewWebsocketHandler(upgrader websocket.Upgrader) *WebsocketHandler {
	return &WebsocketHandler{upgrader}
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

		log.Println("Received message: %s", message)

		responseMsg := fmt.Sprintf("Echo %s", message)

		err = conn.WriteMessage(mt, []byte(responseMsg))

		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}
