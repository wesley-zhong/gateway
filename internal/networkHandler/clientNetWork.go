package networkHandler

import (
	"bytes"
	"encoding/binary"
	"gateway/internal/client"
	"gateway/pkg/log"
	"gateway/pkg/network"
	"time"
)

type ClientNetwork struct {
}

func (clientNetwork *ClientNetwork) OnOpened(c network.ChannelContext) (out []byte, action int) {
	context := client.NewClientContext(c)
	log.Infof("----------  client opened  addr=%s, id=%d", context.Ctx.RemoteAddr(), context.Sid)
	return nil, 0
}

// OnClosed fires when a connection has been closed.
// The parameter err is the last known connection error.
func (clientNetwork *ClientNetwork) OnClosed(c network.ChannelContext, err error) (action int) {
	context := c.Context().(*client.ConnInnerClientContext)
	log.Infof("XXXXXXXXXXXXXXXXXXXX  client closed addr ={} id ={}", c.RemoteAddr(), context)
	return 1

}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (clientNetwork *ClientNetwork) PreWrite(c network.ChannelContext) {
}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (clientNetwork *ClientNetwork) AfterWrite(c network.ChannelContext, b []byte) {

}

// React fires when a socket receives data from the peer.
// Call c.Read() or c.ReadN(n) of Conn c to read incoming data from the peer.
// The parameter out is the return value which is going to be sent back to the peer.
//
// Note that the parameter packet returned from React() is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after React() returns.
// If you have to use packet in a new goroutine, then you need to make a copy of buf and pass this copy
// to that new goroutine.
func (clientNetwork *ClientNetwork) React(packet []byte, c network.ChannelContext) (out []byte, action int) {

	log.Infof("  client receive addr =%s", c.RemoteAddr())

	var msgId int32
	bytebuffer := bytes.NewBuffer(packet)

	binary.Read(bytebuffer, binary.BigEndian, &msgId)
	var length uint32
	binary.Read(bytebuffer, binary.BigEndian, &length)
	log.Infof("---XXXXXXXXXXXXXXXXXXXX ---receive msgId = %d length =%d", msgId, length)
	return nil, 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (clientNetwork *ClientNetwork) Tick() (delay time.Duration, action int) {
	return 1000 * time.Millisecond, 0
}
