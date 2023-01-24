package main

import (
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Client started")
	for {
		conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "wss://stream.binance.com:9443/ws/btcusdt@kline_5m")
		if err != nil {
			fmt.Println("Cannot connect: " + err.Error())
			time.Sleep(time.Duration(5) * time.Second)
			continue
		}
		fmt.Println("Connected to server")
		for i := 0; i < 10; i++ {
			randomNumber := strconv.Itoa(rand.Intn(100))
			msg := []byte(randomNumber)
			err = wsutil.WriteClientMessage(conn, ws.OpText, msg)
			if err != nil {
				fmt.Println("Cannot send: " + err.Error())
				continue
			}
			fmt.Println("Client message send with random number " + randomNumber)
			msg, _, err := wsutil.ReadServerData(conn)
			if err != nil {
				fmt.Println("Cannot receive data: " + err.Error())
				continue
			}
			fmt.Println("Server message received with random number: " + string(msg))
			time.Sleep(time.Duration(5) * time.Second)
		}
		err = conn.Close()
		if err != nil {
			fmt.Println("Cannot close the connection: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Disconnected from server")
	}
}


// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/gobwas/ws"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// type Event struct {
// 	ID       int                    `json:"id"`
// 	Metadata map[string]interface{} `json:"metadata"`
// }

// type MockEventsGenerator struct {
// 	counter      int
// 	tickDuration time.Duration
// }

// func (g *MockEventsGenerator) generate() Event {
// 	g.counter++
// 	return Event{
// 		ID: g.counter,
// 		Metadata: map[string]interface{}{
// 			"type":    "event",
// 			"message": "hello",
// 		},
// 	}
// }

// func (g *MockEventsGenerator) Events(ctx context.Context) <-chan Event {
// 	ch := make(chan Event)

// 	go func() {
// 		defer close(ch)

// 		for {

// 			conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:8080/")
// 			if err != nil {
// 				fmt.Println("Cannot connect: " + err.Error())
// 				time.Sleep(time.Duration(5) * time.Second)
// 				continue
// 			}

			

// 			// conn, _, _, err :=
// 			//  ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:8080/")

// 			// if err != nil {
// 			// 	fmt.Println("Cannot connect: " + err.Error())
// 			// 	continue
// 			// }
// 			// select {
// 			// case <-time.After(g.tickDuration):
// 			// 	ch <- g.generate()
// 			// case <-ctx.Done():
// 			// 	return
// 			// }
// 		}
// 	}()

// 	return ch
// }

// func main() {
// 	r := gin.New()
// 	r.GET("/ws", getWeb)
// 	r.Run("localhost:8080")
// }

// func getWeb(c *gin.Context) {
// 	eventGenerator := MockEventsGenerator{
// 		counter:      0,
// 		tickDuration: time.Second * 2,
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
// 	defer cancel()

// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "the end"))
// 		conn.Close()
// 	}()

// 	for event := range eventGenerator.Events(ctx) {
// 		if err = conn.WriteJSON(event); err != nil {
// 			fmt.Println("failed to write json:", err)
// 			return
// 		}
// 	}
// }
