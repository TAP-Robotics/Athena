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

func (handler *FrameHandler) encodeToJPEG() *gocv.NativeByteBuffer {
	encodedFile, err := gocv.IMEncode(gocv.JPEGFileExt, *handler.Frame)
	if err != nil {
		fmt.Println("Unable to encode frame to jpeg")
	}

	return encodedFile
}

func (handler *FrameHandler) GetJPEGString() []byte {
	encodedFile := handler.encodeToJPEG()
	bytesEnc := encodedFile.GetBytes()
	return bytesEnc
}
