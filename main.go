package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/stream", func(c *gin.Context) {

		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		for {
			conn.WriteMessage(websocket.TextMessage, []byte("PING"))
			time.Sleep(1 * time.Second)
		}
	})

	r.Run() // listen and server on 0.0.0.0:8080
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

