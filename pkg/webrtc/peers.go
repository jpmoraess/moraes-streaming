package webrtc

import (
	"github.com/pion/webrtc/v3"
	"sync"
)

type Peers struct {
	ListLock    sync.RWMutex
	Connections []webrtc.PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrame() {

}
