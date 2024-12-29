package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("Hello World")

	cam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("I love Golang")
	frame := gocv.NewMat()

	for {
		cam.Read(&frame)
		window.IMShow(frame)
		window.WaitKey(1)
	}
}
