package main

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("Hello World")

	cam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("I love Golang")
	frame := gocv.NewMat()

	start := time.Now()

	for {
		timer := time.Now()
		elapsed := timer.Sub(start)
		fmt.Println(elapsed.Milliseconds())
		if elapsed.Milliseconds() >= 100 {
			cam.Read(&frame)
			fmt.Println(frame.Size())
			window.IMShow(frame)
			start = time.Now()
		}
		window.WaitKey(1)
	}
}
