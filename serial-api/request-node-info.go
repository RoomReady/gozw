package serialapi

import (
	"errors"

	"github.com/rmacster/gozw/frame"
	"github.com/rmacster/gozw/protocol"
	"github.com/rmacster/gozw/session"
)

func (s *Layer) RequestNodeInfo(nodeID byte) error {

	done := make(chan *frame.Frame)

	request := &session.Request{
		FunctionID: protocol.FnRequestNodeInfo,
		Payload:    []byte{nodeID},
		HasReturn:  true,
		ReturnCallback: func(err error, ret *frame.Frame) bool {
			done <- ret
			return false
		},
	}

	s.sessionLayer.MakeRequest(request)
	ret := <-done

	if ret == nil {
		return errors.New("Error requesting node information frame")
	}

	status := ret.Payload[1]

	if status == 0 {
		return errors.New("Failed putting node info request in transmit queue")
	}

	return nil
}
