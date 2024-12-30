package frames

import (
	"fmt"

	"gocv.io/x/gocv"
)

type FrameHandler struct {
	opencvInstance *gocv.VideoCapture
	Frame          *gocv.Mat
}

func InitializeFrameHandler() *FrameHandler {
	cvInstance, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println("Cannot open video capture.")
	}
	mat := gocv.NewMat()

	return &FrameHandler{
		opencvInstance: cvInstance,
		Frame:          &mat,
	}
}

func (handler *FrameHandler) GetInstantFrame() {
	handler.opencvInstance.Read(handler.Frame)
}
