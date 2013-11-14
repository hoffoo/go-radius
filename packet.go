package radius

import (
	"net"
)

type Packet struct {
	buff []byte
	read int
	addr *net.UDPAddr
	err  error
}

func CreatePacket(Code, identifier int) {

}

type Code int

const (
	Invalid_Type Code = iota - 1
	Access_Request
	Access_Accept
	Access_Reject
)

const (
	code                int = 0
	identifier              = 1
	len_start               = 2
	len_end                 = 3
	authenticator_start     = 4
	authenticator_end       = 20
)

func (p *Packet) Code() Code {
	pt := packetType(p.buff[0])
	return pt
}

func (p *Packet) Identifier() int {
	return int(p.buff[1])
}

func (p *Packet) Len() int {
	return int(p.buff[2])
}

func packetType(code byte) Code {
	switch Code(code) {
	case Access_Request:
		return Access_Request
	case Access_Accept:
		return Access_Accept
	case Access_Reject:
		return Access_Reject
	default:
		return Invalid_Type
	}
}
