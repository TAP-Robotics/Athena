package main

import (
	"fmt"
	"oddysseus/internal/frames"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("Hello World")

	frame_handler := frames.InitializeFrameHandler()

	window := gocv.NewWindow("I love Golang")
	start := time.Now()

	for {
		timer := time.Now()
		elapsed := timer.Sub(start)
		fmt.Println(elapsed.Milliseconds())
		if elapsed.Milliseconds() >= 100 {
			frame_handler.GetInstantFrame()
			window.IMShow(*frame_handler.Frame)
			start = time.Now()
		}
		window.WaitKey(1)
	}
}
