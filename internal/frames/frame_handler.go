package frames

import (
	"fmt"

	"github.com/rs/zerolog/log"
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

func (handler *FrameHandler) encodeToImage() *gocv.NativeByteBuffer {
	quality := []int{gocv.IMWriteWebpQuality, 10}
	encodedFile, err := gocv.IMEncodeWithParams(".webp", *handler.Frame, quality)
	if err != nil {
		fmt.Println("Unable to encode frame to jpeg")
	}

	return encodedFile
}

func (handler *FrameHandler) GetImageString() []byte {
	encodedFile := handler.encodeToImage()
	bytesEnc := encodedFile.GetBytes()

	sInB := len(bytesEnc)
	sInMb := float64(sInB) / (1024 * 1024)

	log.Debug().Float64("size_in_mb", sInMb).Msg("Encoded sizm")

	return bytesEnc
}
