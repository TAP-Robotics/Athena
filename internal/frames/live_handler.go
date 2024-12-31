package frames

import (
	"encoding/json"
	"oddysseus/internal/shared"
	"time"

	"github.com/go-zeromq/zmq4"
	"github.com/rs/zerolog/log"
	"gocv.io/x/gocv"
)

type LiveHandler struct {
	frameHandler *FrameHandler
	socket       zmq4.Socket
}

func NewVisionHandler(fHand *FrameHandler, sock zmq4.Socket) *LiveHandler {
	return &LiveHandler{
		frameHandler: fHand,
		socket:       sock,
	}
}

func (lHand *LiveHandler) HandleDealer() {
	window := gocv.NewWindow("Client / Robot's Vision (GoLang)")

	// start a timer, will be used to check if 100ms has passed.
	start := time.Now()

	for {
		timer := time.Now()
		elapsed := timer.Sub(start)        // get the difference between start and end (timer)
		if elapsed.Milliseconds() >= 140 { // 140 ms is the average inference time of the server in this environment (around 7.14 F/S).
			lHand.frameHandler.GetInstantFrame()
			stringEncodedImageCapture := lHand.frameHandler.GetPNGString()

			toSend := &shared.SendTempalte{
				Message: "vision_infer",
				Content: &stringEncodedImageCapture,
			}

			marshaledData, err := json.Marshal(toSend)
			if err != nil {
				log.Warn().Msg("Marshaling send template failed: " + err.Error())
				return
			}

			message := zmq4.NewMsgString(string(marshaledData))
			if err := lHand.socket.Send(message); err != nil {
				log.Error().Msg("Vision/Frame send failed: " + err.Error())
				return
			}

			window.IMShow(*lHand.frameHandler.Frame)
			window.WaitKey(1)

			start = time.Now() // restart timer by setting start time as now
		}
	}

}
