package main

import (
	"context"
	"encoding/json"
	"fmt"
	"oddysseus/internal/frames"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"gocv.io/x/gocv"
)

type SendTempalte struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func main() {
	fmt.Println("Hello World")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	connection, _, err := websocket.Dial(ctx, "ws://localhost:7207", nil)
	if err != nil {
		fmt.Println("Can't establish websocket connection, exiting...")
		panic(err)
	}
	defer connection.CloseNow()

	frame_handler := frames.InitializeFrameHandler()

	window := gocv.NewWindow("I love Golang")
	start := time.Now()

	for {
		timer := time.Now()
		elapsed := timer.Sub(start)
		if elapsed.Milliseconds() >= 140 {
			frame_handler.GetInstantFrame()
			stringedBytes := frame_handler.GetJPEGString()

			toSend := &SendTempalte{
				Message: "infer",
				Content: &stringedBytes,
			}
			marshaled, err := json.Marshal(toSend)
			if err != nil {
				fmt.Println("Failed to marshal frame toSend data.")
			}

			err = wsjson.Write(ctx, connection, string(marshaled))
			if err != nil {
				fmt.Println("bruh error")
			}

			window.IMShow(*frame_handler.Frame)
			start = time.Now()
		}
		window.WaitKey(1)
	}
}
